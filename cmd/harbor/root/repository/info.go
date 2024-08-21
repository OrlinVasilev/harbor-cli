// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package repository

import (
	"github.com/goharbor/harbor-cli/pkg/api"
	"github.com/goharbor/harbor-cli/pkg/prompt"
	"github.com/goharbor/harbor-cli/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func RepoInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "info",
		Short:   "Get repository information",
		Example: `  harbor repo info <project_name>/<repo_name>`,
		Long:    `Get information of a particular repository in a project`,
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			if len(args) > 0 {
				projectName, repoName := utils.ParseProjectRepo(args[0])
				err = api.RepoInfo(projectName, repoName)
			} else {
				projectName := prompt.GetProjectNameFromUser()
				repoName := prompt.GetRepoNameFromUser(projectName)
				err = api.RepoInfo(projectName, repoName)
			}
			if err != nil {
				log.Errorf("failed to get repository information: %v", err)
			}

		},
	}

	return cmd
}
