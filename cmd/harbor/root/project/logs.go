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

			FormatFlag := viper.GetString("output-format")
			if FormatFlag != "" {
				if FormatFlag == "json" {
					utils.PrintPayloadInJSONFormat(resp)
					return
				}
				if FormatFlag == "yaml" {
					utils.PrintPayloadInYAMLFormat(resp)
					return
				}
				log.Errorf("Unable to output in the specified '%s' format", FormatFlag)
				return
			}
			auditLog.LogsProject(resp.Payload)

		},
	}

	return cmd
}
