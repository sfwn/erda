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

package migration

import (
	"errors"
	"fmt"
	"strings"

	"github.com/erda-project/erda/modules/gittar/models"
)

func NewAuth() error {
	db, err := models.OpenDB()
	defer db.Close()
	if err != nil {
		return err
	}
	var repos []models.Repo
	db.Raw(`
SELECT
dice_app.id as app_id,
dice_app.name as app_name,
org_id,project_id,project_name,
dice_org.name as org_name,
dice_app.git_repo_abbrev as path
from dice_app
INNER JOIN dice_org on dice_app.org_id = dice_org.id
`).Scan(&repos)

	errorRepos := []string{}
	tx := db.Begin()
	for _, repo := range repos {
		var currentRepo models.Repo
		err = db.Where("org_name=? and project_name=? and app_name=? ", repo.OrgName, repo.ProjectName, repo.AppName).
			First(&currentRepo).Error

		if err != nil {
			//无法查到报err,说明还没未执行过
			splits := strings.Split(repo.Path, "/")
			//无法解析的跳过
			if len(splits) < 2 {
				errorRepos = append(errorRepos, fmt.Sprintf("%v:%v", repo.AppID, repo.Path))
				continue
			}
			oldOrg := splits[0]
			oldRepo := splits[1]
			err := db.Create(&repo).Error
			if err != nil {
				tx.Rollback()
				return err
			}
			err = db.Exec("update gittar_web_hooks set repo_id = ? where org = ? and repo =? ", repo.ID, oldOrg, oldRepo).Error
			if err != nil {
				tx.Rollback()
				return err
			}
			err = db.Exec("update gittar_merge_requests set repo_id = ? where org = ? and repo =? ", repo.ID, oldOrg, oldRepo).Error
			if err != nil {
				tx.Rollback()
				return err
			}
			err = db.Exec("update gittar_notes set repo_id = ? where org = ? and repo =? ", repo.ID, oldOrg, oldRepo).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		} else {
			//已经存在
			tx.Rollback()
			return errors.New("new_auth already migrated")
		}
	}
	tx.Commit()

	if len(errorRepos) > 0 {
		return errors.New("failed repos:" + strings.Join(errorRepos, ","))
	}
	return nil
}
