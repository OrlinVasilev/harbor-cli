package registry

import (
	"context"
	"fmt"
	"strconv"

	"github.com/akshatdalton/harbor-cli/cmd/utils"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/registry"
	"github.com/spf13/cobra"
)

type getRegistryOptions struct {
	id int64
}

// NewGetRegistryCommand creates a new `harbor get registry` command
func NewGetRegistryCommand() *cobra.Command {
	var opts getRegistryOptions

	cmd := &cobra.Command{
		Use:   "registry [ID]",
		Short: "get registry by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Printf("Invalid argument: %s. Expected an integer.\n", args[0])
				return err
			}
			opts.id = int64(id)
			return runGetRegistry(opts)
		},
	}

	return cmd
}

func runGetRegistry(opts getRegistryOptions) error {
	client := utils.GetClient(nil)
	ctx := context.Background()
	response, err := client.Registry.GetRegistry(ctx, &registry.GetRegistryParams{ID: opts.id})

	if err != nil {
		return err
	}

	utils.PrintPayloadInJSONFormat(response.GetPayload())
	return nil
}
