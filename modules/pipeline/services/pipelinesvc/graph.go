package pipelinesvc

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/pipeline/services/apierrors"
	"terminus.io/dice/dice/pkg/parser/pipelineyml"
)

func (s *PipelineSvc) PipelineYmlGraph(req *apistructs.PipelineYmlParseGraphRequest) (*apistructs.PipelineYml, error) {
	graph, err := pipelineyml.ConvertToGraphPipelineYml([]byte(req.PipelineYmlContent))
	if err != nil {
		return nil, apierrors.ErrParsePipelineYml.InvalidParameter(err)
	}

	if graph == nil {
		return graph, nil
	}

	// 是否需要展平 snippet
	if req.NeedFlatSnippet {
		caches, err := loadSnippetCaches(graph, req.SnippetConfig, req.GlobalSnippetConfigLabels)
		if err != nil {
			return nil, err
		}
		// 设置层级结构的stage
		s.PipelineNestedSnippetStages(graph, req.GlobalSnippetConfigLabels, caches.PipelineYamlCaches)
		// 设置展平结构的stage
		if err := s.flatSnippet(graph, req.GlobalSnippetConfigLabels, caches.PipelineYamlCaches); err != nil {
			return nil, err
		}
	}

	// 设置logo和名称
	s.loadGraphActionNameAndLogo(graph)

	return graph, nil
}

// 加载缓存
func loadSnippetCaches(graph *apistructs.PipelineYml, snippetConfig *apistructs.SnippetConfig, globalSnippetConfigLabels map[string]string) (*pipelineyml.SnippetCache, error) {
	// 初始化缓存
	cache := pipelineyml.SnippetCache{}
	if snippetConfig != nil {
		cache.PipelineYamlCaches = append(cache.PipelineYamlCaches, pipelineyml.SnippetPipelineYmlCache{
			SnippetConfig: pipelineyml.SnippetConfig{
				Name:   snippetConfig.Name,
				Source: snippetConfig.Source,
				Labels: snippetConfig.Labels,
			},
			PipelineYaml: graph,
		})
	}
	err := cache.InitCache(graph.Stages, globalSnippetConfigLabels)
	if err != nil {
		return nil, apierrors.ErrParsePipelineYml.InvalidParameter(err)
	}

	return &cache, nil
}

// 将其中的 action 全部展开到一个数组中
func (s *PipelineSvc) flatSnippet(graph *apistructs.PipelineYml, globalSnippetConfigLabels map[string]string, pipelineYamlCaches []pipelineyml.SnippetPipelineYmlCache) error {
	// 设置展平的结构的stage
	flatActions, err := s.PipelineYmlFlatGraph(graph, globalSnippetConfigLabels, pipelineYamlCaches)
	if err != nil {
		return apierrors.ErrParsePipelineYml.InvalidParameter(err)
	}
	graph.FlatActions = flatActions
	return nil
}

func (s *PipelineSvc) loadGraphActionNameAndLogo(graph *apistructs.PipelineYml) {
	var extensionSearchRequest = apistructs.ExtensionSearchRequest{}
	extensionSearchRequest.YamlFormat = true
	for _, stage := range graph.Stages {
		for _, action := range stage {
			extensionSearchRequest.Extensions = append(extensionSearchRequest.Extensions, action.Type)
		}
	}

	resultMap, err := s.bdl.SearchExtensions(extensionSearchRequest)
	if err != nil {
		logrus.Errorf("pipelineYmlGraph to SearchExtensions error: %v", err)
		return
	}
	for _, stage := range graph.Stages {
		for _, action := range stage {
			if pipelineyml.IsSnippetType(pipelineyml.ActionType(action.Type)) {
				action.LogoUrl = pipelineyml.SnippetLogo
				action.DisplayName = pipelineyml.SnippetDisplayName
				action.Description = pipelineyml.SnippetDesc
				continue
			}

			version, ok := resultMap[action.Type]
			if !ok {
				continue
			}

			specYmlStr, ok := version.Spec.(string)
			if !ok {
				continue
			}

			var actionSpec apistructs.ActionSpec
			if err := yaml.Unmarshal([]byte(specYmlStr), &actionSpec); err != nil {
				logrus.Errorf("pipelineYmlGraph Unmarshal spec error: %v", err)
				continue
			}

			action.DisplayName = actionSpec.DisplayName
			action.LogoUrl = actionSpec.LogoUrl
			action.Description = actionSpec.Desc
		}
	}
}

// snippet 中的 action 又会包含 snippet 结构，这样前端可以弄成一个拓扑图
func (s *PipelineSvc) PipelineNestedSnippetStages(graph *apistructs.PipelineYml, globalSnippetConfigLabels map[string]string, pipelineYamlCaches []pipelineyml.SnippetPipelineYmlCache) {
	if graph == nil {
		return
	}

	if globalSnippetConfigLabels == nil {
		globalSnippetConfigLabels = map[string]string{}
	}

	stages := graph.Stages
	if stages == nil {
		return
	}

	doPipelineNestedSnippetStages(stages, globalSnippetConfigLabels, pipelineYamlCaches)
}

// 递归设置snippet的graph
func doPipelineNestedSnippetStages(stages [][]*apistructs.PipelineYmlAction,
	globalSnippetConfigLabels map[string]string,
	pipelineYamlCaches []pipelineyml.SnippetPipelineYmlCache) {

	for stageIndex, stage := range stages {
		for actionIndex, action := range stage {
			if pipelineyml.IsSnippetType(pipelineyml.ActionType(action.Type)) {
				snippetConfig := pipelineyml.HandleSnippetConfigLabel(&pipelineyml.SnippetConfig{
					Source: action.SnippetConfig.Source,
					Name:   action.SnippetConfig.Name,
					Labels: action.SnippetConfig.Labels,
				}, globalSnippetConfigLabels)

				var graph = pipelineyml.GetCacheYaml(pipelineYamlCaches, &snippetConfig)

				if graph == nil {
					continue
				}
				doPipelineNestedSnippetStages(graph.Stages, globalSnippetConfigLabels, pipelineYamlCaches)

				var snippetStages = apistructs.SnippetStages{
					Params:  graph.Params,
					Outputs: graph.Outputs,
					Stages:  graph.Stages,
				}
				stages[stageIndex][actionIndex].SnippetStages = &snippetStages
			}
		}
	}
}

func (s *PipelineSvc) PipelineYmlFlatGraph(graph *apistructs.PipelineYml, globalSnippetConfigLabels map[string]string, pipelineYamlCaches []pipelineyml.SnippetPipelineYmlCache) ([]*apistructs.PipelineYmlAction, error) {
	if graph == nil {
		return nil, nil
	}

	if globalSnippetConfigLabels == nil {
		globalSnippetConfigLabels = map[string]string{}
	}

	stages := graph.Stages
	if stages == nil {
		return nil, nil
	}

	actions, err := setFlatSnippetAction(0, stages, pipelineYamlCaches, globalSnippetConfigLabels, "")
	if err != nil {
		return nil, err
	}

	return actions, nil
}

func setFlatSnippetAction(myDeep int, stages [][]*apistructs.PipelineYmlAction,
	pipelineYamlCaches []pipelineyml.SnippetPipelineYmlCache,
	globalSnippetConfigLabels map[string]string,
	preAlias string) ([]*apistructs.PipelineYmlAction, error) {

	if stages == nil {
		return nil, nil
	}

	// 判定深度是否大于9，大于9下面+1深度就大于10了
	if myDeep > pipelineyml.MaxDeep {
		return nil, fmt.Errorf("getSnippetPipelineYaml error: snippet deep > %v，Please verify whether there is a circular dependency", pipelineyml.MaxDeep)
	}

	// 传递给下一个递归的深度
	deep := myDeep + 1

	var allAction []*apistructs.PipelineYmlAction

	for _, stage := range stages {

		var stageAllAction []*apistructs.PipelineYmlAction

		for _, action := range stage {

			if pipelineyml.IsSnippetType(pipelineyml.ActionType(action.Type)) {

				snippetConfig := pipelineyml.HandleSnippetConfigLabel(&pipelineyml.SnippetConfig{
					Source: action.SnippetConfig.Source,
					Name:   action.SnippetConfig.Name,
					Labels: action.SnippetConfig.Labels,
				}, globalSnippetConfigLabels)

				var graph = pipelineyml.GetCacheYaml(pipelineYamlCaches, &snippetConfig)
				if graph != nil {
					var preName = action.Alias
					if preAlias != "" {
						preName = preAlias + pipelineyml.SnippetActionNameLinkAddr + preName
					}

					list, err := setFlatSnippetAction(deep, graph.Stages, pipelineYamlCaches, globalSnippetConfigLabels, preName)
					if err != nil {
						return nil, err
					}
					if list != nil {
						stageAllAction = append(stageAllAction, list...)
					}
				}
			} else {
				if preAlias != "" {
					action.Alias = preAlias + pipelineyml.SnippetActionNameLinkAddr + action.Alias
				}
				stageAllAction = append(stageAllAction, action)
			}
		}

		allAction = append(allAction, stageAllAction...)
	}

	return allAction, nil
}
