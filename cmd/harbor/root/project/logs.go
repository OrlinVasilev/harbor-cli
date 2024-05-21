package project

import (
	"context"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
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
				resp, err = runLogsProject(args[0])
			} else {
				projectName := utils.GetProjectNameFromUser()
				resp, err = runLogsProject(projectName)
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

func runLogsProject(projectName string) (*project.GetLogsOK, error) {
	credentialName := viper.GetString("current-credential-name")
	client := utils.GetClientByCredentialName(credentialName)
	ctx := context.Background()
	response, err := client.Project.GetLogs(ctx, &project.GetLogsParams{
		ProjectName: projectName,
		Context:     ctx,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}
