package user

import (
	"strconv"

	"github.com/goharbor/harbor-cli/pkg/api"
	"github.com/goharbor/harbor-cli/pkg/prompt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func ElevateUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "elevate",
		Short: "elevate user",
		Long:  "elevate user to admin role",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			var userId int64
			if len(args) > 0 {
				userId, _ = strconv.ParseInt(args[0], 10, 64)

			} else {
				userId = prompt.GetUserIdFromUser()
			}

			// Todo : Ask for the confirmation before elevating the user to admin role

			err = api.ElevateUser(userId)

			if err != nil {
				log.Errorf("failed to elevate user: %v", err)
			}

		},
	}

	return cmd
}
