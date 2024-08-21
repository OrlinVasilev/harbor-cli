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

package user

import (
	"github.com/goharbor/harbor-cli/pkg/api"
	"github.com/goharbor/harbor-cli/pkg/utils"
	"github.com/goharbor/harbor-cli/pkg/views/user/list"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func UserListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "list users",
		Args:    cobra.NoArgs,
		Aliases: []string{"ls"},
		Run: func(cmd *cobra.Command, args []string) {
			response, err := api.ListUsers()
			if err != nil {
				log.Errorf("failed to list users: %v", err)
				return
			}
			FormatFlag := viper.GetString("output-format")
			if FormatFlag != "" {
				utils.PrintPayloadInJSONFormat(response.Payload)
			} else {
				list.ListUsers(response.Payload)
			}
		},
	}

	return cmd

}
