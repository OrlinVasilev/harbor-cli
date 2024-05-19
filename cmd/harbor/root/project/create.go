package project

import (
	"context"
	"strconv"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
	"github.com/goharbor/harbor-cli/pkg/utils"
	"github.com/goharbor/harbor-cli/pkg/views/project/create"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// CreateProjectCommand creates a new `harbor create project` command
func CreateProjectCommand() *cobra.Command {
	var opts create.CreateView

	cmd := &cobra.Command{
		Use:   "create",
		Short: "create project",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			createView := &create.CreateView{
				ProjectName:  opts.ProjectName,
				Public:       opts.Public,
				RegistryID:   opts.RegistryID,
				StorageLimit: opts.StorageLimit,
			}
			credentialName := viper.GetString("current-credential-name")
			client := utils.GetClientByCredentialName(credentialName)
			ctx := context.Background()
			if opts.ProjectName != "" && opts.RegistryID != "" && opts.StorageLimit != "" {
				err = RunCreateProject(ctx, opts, client.Project)
			} else {
				createView = createProjectView(createView)
				err = RunCreateProject(ctx, *createView, client.Project)
			}

			if err != nil {
				log.Errorf("failed to create project: %v", err)
			}

		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.ProjectName, "name", "", "", "Name of the project")
	flags.BoolVarP(&opts.Public, "public", "", true, "Project is public or private")
	flags.StringVarP(&opts.RegistryID, "registry-id", "", "", "ID of referenced registry when creating the proxy cache project")
	flags.StringVarP(&opts.StorageLimit, "storage-limit", "", "-1", "Storage quota of the project")

	return cmd
}

func createProjectView(createView *create.CreateView) *create.CreateView {
	if createView == nil {
		createView = &create.CreateView{
			ProjectName:  "",
			Public:       true,
			RegistryID:   "",
			StorageLimit: "-1",
		}
	}

	create.CreateProjectView(createView)

	return createView

}

func RunCreateProject(ctx context.Context, opts create.CreateView, projectInterface ProjectInterface) error {

	public := strconv.FormatBool(opts.Public)

	response, err := projectInterface.CreateProject(ctx, &project.CreateProjectParams{Project: &models.ProjectReq{ProjectName: opts.ProjectName, Public: &opts.Public, Metadata: &models.ProjectMetadata{Public: public}}})

	if err != nil {
		return err
	}

	if response != nil {
		log.Info("Project created successfully")
	}
	return nil
}
