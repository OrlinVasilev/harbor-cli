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

package project

import (
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
	"github.com/goharbor/harbor-cli/pkg/api"
	"github.com/goharbor/harbor-cli/pkg/prompt"
	"github.com/goharbor/harbor-cli/pkg/utils"
	auditLog "github.com/goharbor/harbor-cli/pkg/views/project/logs"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func LogsProjectCommmand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logs",
		Short: "get project logs",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			var resp *project.GetLogsOK
			if len(args) > 0 {
				resp, err = api.LogsProject(args[0])
			} else {
				projectName := prompt.GetProjectNameFromUser()
				resp, err = api.LogsProject(projectName)
			}

			if err != nil {
				log.Fatalf("failed to get project logs: %v", err)
			}
			auditLog.LogsProject(resp.Payload)

			FormatFlag := viper.GetString("output-format")
			if FormatFlag != "" {
				utils.PrintPayloadInJSONFormat(resp)
				return
			}

		},
	}

	return cmd
}
