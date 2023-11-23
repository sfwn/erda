// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cr

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/mohae/deepcopy"
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
	"github.com/sirupsen/logrus"
	"sigs.k8s.io/yaml"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/internal/tools/gittar/ai/cr/aiutil"
	"github.com/erda-project/erda/internal/tools/gittar/models"
	"github.com/erda-project/erda/internal/tools/gittar/pkg/gitmodule"
)

type OneChangedFile struct {
	FileName     string
	CodeLanguage string
	Truncated    bool
	CodeSnippets []CodeSnippet

	mr       *apistructs.MergeRequestInfo
	diffFile *gitmodule.DiffFile
	user     *models.User
}

func newFileReviewer(diffFile *gitmodule.DiffFile, user *models.User, mr *apistructs.MergeRequestInfo) FileCodeReviewer {
	fr := OneChangedFile{diffFile: diffFile, user: user, mr: mr}
	fr.FileName = diffFile.Name
	fr.CodeLanguage = strings.TrimPrefix(filepath.Ext(diffFile.Name), ".")
	return &fr
}

func (r *OneChangedFile) GetFileName() string {
	return r.FileName
}

// CodeReview for file level, invoke once with all code snippets.
func (r *OneChangedFile) CodeReview() string {
	if r.diffFile == nil {
		return ""
	}
	r.parseCodeSnippets()

	// invoke ai
	result := aiutil.InvokeAI(r.ConstructAIRequest(), r.user)
	if result == "" {
		return ""
	}

	// handle response
	var res FileReviewResult
	if err := json.Unmarshal([]byte(result), &res); err != nil {
		logrus.Warnf("failed to unmarshal ai result, err: %s", err)
	}
	// group result by snippet index
	snippetIndexIssues := make(map[int][]FileReviewResultItem)
	for _, item := range res.Result {
		if item.SnippetIndex >= len(r.CodeSnippets) {
			continue
		}
		snippetIndexIssues[item.SnippetIndex] = append(snippetIndexIssues[item.SnippetIndex], item)
	}
	// handle each snippet index
	var s string
	for snippetIndex := range r.CodeSnippets {
		// add original code
		s += fmt.Sprintf("**Snippet:**\n%s\n", r.CodeSnippets[snippetIndex].getMarkdownCode())
		for _, issue := range snippetIndexIssues[snippetIndex] {
			s += fmt.Sprintf("**RiskLevel:** %s\n", issue.RiskLevel)
			s += fmt.Sprintf("\n%s\n\n", issue.Details)
		}
		s += "\n"
	}
	return s
}

type (
	FileReviewResult struct {
		Result []FileReviewResultItem `json:"fileReviewResult,omitempty"`
	}
	FileReviewResultItem struct {
		SnippetIndex int    `json:"snippetIndex,omitempty"`
		RiskLevel    string `json:"riskLevel,omitempty"`
		Details      string `json:"details,omitempty"`
	}
)

//go:embed prompt.yaml
var promptYaml string

type PromptStruct = struct {
	Messages []openai.ChatCompletionMessage `yaml:"messages"`
}

var promptStruct PromptStruct

func init() {
	if err := yaml.Unmarshal([]byte(promptYaml), &promptStruct); err != nil {
		panic(err)
	}
}

func (r *OneChangedFile) ConstructAIRequest() openai.ChatCompletionRequest {
	msgs := deepcopy.Copy(promptStruct.Messages).([]openai.ChatCompletionMessage)

	req := openai.ChatCompletionRequest{
		Messages: msgs,
		Stream:   false,
		Functions: []openai.FunctionDefinition{
			{
				Name:        "create-cr-note",
				Description: "create code review note",
				Parameters: &jsonschema.Definition{
					Type:        jsonschema.Object,
					Description: "create code review note for each file",
					Properties: map[string]jsonschema.Definition{
						"fileReviewResult": {
							Type:        jsonschema.Array,
							Description: "review result for each file",
							Items: &jsonschema.Definition{
								Type: jsonschema.Object,
								Properties: map[string]jsonschema.Definition{
									"snippetIndex": {
										Type:        jsonschema.Integer,
										Description: "snippet index",
									},
									"riskLevel": {
										Type:        jsonschema.String,
										Description: "risk level (使用普通人能听懂的话，例如：建议、严重、致命等)",
									},
									"details": {
										Type:        jsonschema.String,
										Description: "details, one or array. (总结这段代码；对于每个 issue，如果可以，请给出用于修复的示例代码)",
									},
								},
								Required: []string{"snippetIndex", "riskLevel", "details"},
							},
							Required: []string{"fileReviewResult"},
						},
					},
					Items: &jsonschema.Definition{},
				},
			},
		},
		FunctionCall: openai.FunctionCall{
			Name: "create-cr-note",
		},
	}

	var tmplArgs struct {
		CodeLanguage string
		FileName     string
		FileContents string
	}
	tmplArgs.CodeLanguage = r.CodeLanguage
	tmplArgs.FileName = r.FileName

	type SnippetContent struct {
		SnippetIndex  int    `json:"snippetIndex"`
		PromptContent string `json:"promptContent"`
	}
	var changedFileContents []string
	for i, cs := range r.CodeSnippets {
		sc := SnippetContent{SnippetIndex: i, PromptContent: cs.SelectedCode}
		b, _ := json.Marshal(sc)
		changedFileContents = append(changedFileContents, string(b))
	}
	tmplArgs.FileContents = strings.Join(changedFileContents, "\n\n")

	for i := range req.Messages {
		t, _ := template.New("").Parse(req.Messages[i].Content)
		if t == nil {
			continue
		}
		buffer := bytes.NewBuffer(nil)
		_ = t.Execute(buffer, tmplArgs)
		req.Messages[i].Content = buffer.String()
	}

	return req
}

func (r *OneChangedFile) parseCodeSnippets() {
	for _, section := range r.diffFile.Sections {
		var changes []string
		truncated := false
		for _, line := range section.Lines {
			s := line.Content
			switch line.Type {
			case gitmodule.DIFF_LINE_ADD:
				s = "+" + s
			case gitmodule.DIFF_LINE_DEL:
				s = "-" + s
			case "": // ignore: \ No newline at end of file
				continue
			case gitmodule.DIFF_LINE_SECTION:
			default:
				s = " " + s
			}
			if s == "" {
				continue
			}
			changes = append(changes, s)
		}
		if len(changes) == 0 {
			continue
		}
		if len(changes) > MAX_FILE_CHANGES_CHAR_SIZE {
			changes = changes[:MAX_FILE_CHANGES_CHAR_SIZE]
			truncated = true
			r.Truncated = true
		}
		codeSnippet := CodeSnippet{
			CodeLanguage: strings.TrimPrefix(filepath.Ext(r.diffFile.Name), "."),
			SelectedCode: strings.Join(changes, "\n"),
			Truncated:    truncated,
			user:         r.user,
		}
		r.CodeSnippets = append(r.CodeSnippets, codeSnippet)
	}
}
