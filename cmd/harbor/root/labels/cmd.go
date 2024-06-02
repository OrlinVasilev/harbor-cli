package labels

import "github.com/spf13/cobra"

func Labels() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "label",
		Short: "Manage labels in Harbor",
	}

	return cmd
}
