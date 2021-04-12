// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package extmarketsvc

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/pkg/parser/diceyml"
	"github.com/erda-project/erda/pkg/parser/pipelineyml"
)

type SearchOption struct {
	NeedRender   bool
	Placeholders map[string]string
}
type OpOption func(*SearchOption)

func SearchActionWithRender(placeholders map[string]string) OpOption {
	return func(so *SearchOption) {
		so.NeedRender = true
		so.Placeholders = placeholders
	}
}

// doSearchRemoteActions search actions from dicehub.
func (s *ExtMarketSvc) doSearchRemoteActions(items []string) (map[string]apistructs.ExtensionVersion, error) {
	req := apistructs.ExtensionSearchRequest{Extensions: items, YamlFormat: true}
	return s.bdl.SearchExtensions(req)
}

// SearchActions try to get from cache.
// each search item: ActionType 或 ActionType@version
// return: allDiceYml, allSpec, error
func (s *ExtMarketSvc) SearchActions(items []string, ops ...OpOption) (map[string]*diceyml.Job, map[string]*apistructs.ActionSpec, error) {
	// get from cache first
	var itemsNeedSearch []string
	cachedActionsByItem := make(map[string]apistructs.ExtensionVersion)
	for _, item := range items {
		ext := extCaches.getExt(item)
		if ext == nil {
			// not found in cache, need search later
			itemsNeedSearch = append(itemsNeedSearch, item)
			continue
		}
		// found in cache
		cachedActionsByItem[item] = *ext
	}

	// search from dicehub
	searchedActionsByItem, err := s.doSearchRemoteActions(itemsNeedSearch)
	if err != nil {
		return nil, nil, err
	}
	// update to cache
	for item, ext := range searchedActionsByItem {
		extCaches.updateExt(item, ext)
	}

	// merge actions
	allActionsByItem := make(map[string]apistructs.ExtensionVersion)
	for item, ext := range cachedActionsByItem {
		allActionsByItem[item] = ext
	}
	for item, ext := range searchedActionsByItem {
		allActionsByItem[item] = ext
	}

	// search
	so := SearchOption{NeedRender: false, Placeholders: nil}
	for _, op := range ops {
		op(&so)
	}
	actionDiceYmlJobMap := make(map[string]*diceyml.Job)
	for nameVersion, action := range allActionsByItem {
		if action.NotExist() {
			errMsg := fmt.Sprintf("action %q not exist in Extension Market", nameVersion)
			logrus.Errorf("[alert] %s", errMsg)
			return nil, nil, errors.New(errMsg)
		}
		diceYmlStr, ok := action.Dice.(string)
		if !ok {
			errMsg := fmt.Sprintf("failed to search action from extension market, action: %s, err: %s", nameVersion, "action's dice.yml is not string")
			logrus.Errorf("[alert] %s, action's dice.yml: %#v", errMsg, action.Dice)
			return nil, nil, errors.New(errMsg)
		}
		if so.NeedRender && len(so.Placeholders) > 0 {
			rendered, err := pipelineyml.RenderSecrets([]byte(diceYmlStr), so.Placeholders)
			if err != nil {
				errMsg := fmt.Sprintf("failed to render action's dice.yml, action: %s, err: %v", nameVersion, err)
				logrus.Errorf("[alert] %s, action's dice.yml: %#v", errMsg, action.Dice)
				return nil, nil, errors.New(errMsg)
			}
			diceYmlStr = string(rendered)
		}
		diceYml, err := diceyml.New([]byte(diceYmlStr), false)
		if err != nil {
			errMsg := fmt.Sprintf("failed to parse action's dice.yml, action: %s, err: %v", nameVersion, err)
			logrus.Errorf("[alert] %s, action's dice.yml: %#v", errMsg, action.Dice)
			return nil, nil, errors.New(errMsg)
		}
		for _, job := range diceYml.Obj().Jobs {
			actionDiceYmlJobMap[nameVersion] = job
			break
		}
	}
	actionSpecMap := make(map[string]*apistructs.ActionSpec)
	for nameVersion, action := range allActionsByItem {
		actionSpecMap[nameVersion] = nil
		specYmlStr, ok := action.Spec.(string)
		if !ok {
			errMsg := fmt.Sprintf("failed to search action from extension market, action: %s, err: %s", nameVersion, "action's spec.yml is not string")
			logrus.Errorf("[alert] %s, action's spec.yml: %#v", errMsg, action.Spec)
			return nil, nil, errors.New(errMsg)
		}
		var actionSpec apistructs.ActionSpec
		if err := yaml.Unmarshal([]byte(specYmlStr), &actionSpec); err != nil {
			errMsg := fmt.Sprintf("failed to parse action's spec.yml, action: %s, err: %v", nameVersion, err)
			logrus.Errorf("[alert] %s, action's spec.yml: %#v", errMsg, action.Spec)
			return nil, nil, errors.New(errMsg)
		}
		actionSpecMap[nameVersion] = &actionSpec
	}

	return actionDiceYmlJobMap, actionSpecMap, nil
}
