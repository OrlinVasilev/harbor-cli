package project

import (
	proj "github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
	"github.com/goharbor/harbor-cli/pkg/api"
	"github.com/goharbor/harbor-cli/pkg/prompt"
	"github.com/goharbor/harbor-cli/pkg/utils"
	auditLog "github.com/goharbor/harbor-cli/pkg/views/project/logs"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func LogsProjectCommmand() *cobra.Command {
	var opts api.ListFlags

	cmd := &cobra.Command{
		Use:   "logs",
		Short: "get project logs",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			var resp *proj.GetLogsOK
			if len(args) > 0 {
				resp, err = api.LogsProject(args[0])
			} else {
				projectName := prompt.GetProjectNameFromUser()
				resp, err = api.LogsProject(projectName)
			}

			if err != nil {
				log.Fatalf("failed to get project logs: %v", err)
				return
			}

			FormatFlag := viper.GetString("output-format")
			if FormatFlag != "" {
				err = utils.PrintFormat(resp, FormatFlag)
				if err != nil {
					log.Error(err)
				}
			} else {
				auditLog.LogsProject(resp.Payload)
			}
		},
	}

	flags := cmd.Flags()
	flags.Int64VarP(&opts.Page, "page", "", 1, "Page number")
	flags.Int64VarP(&opts.PageSize, "page-size", "", 10, "Size of per page")
	flags.StringVarP(&opts.Q, "query", "q", "", "Query string to query resources")
	flags.StringVarP(&opts.Sort, "sort", "", "", "Sort the resource list in ascending or descending order")

	return cmd
}
