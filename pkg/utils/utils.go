package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/user"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Returns Harbor v2 client for given clientConfig

func PrintPayloadInJSONFormat(payload any) {
	if payload == nil {
		return
	}

	jsonStr, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonStr))
}

func ParseProjectRepo(projectRepo string) (string, string) {
	split := strings.Split(projectRepo, "/")
	if len(split) != 2 {
		log.Fatalf("invalid project/repository format: %s", projectRepo)
	}
	return split[0], split[1]
}

func GetUserIdFromUser() int64 {
	userId := make(chan int64)

	go func() {
		credentialName := viper.GetString("current-credential-name")
		client := GetClientByCredentialName(credentialName)
		ctx := context.Background()
		response, err := client.User.ListUsers(ctx, &user.ListUsersParams{})
		if err != nil {
			log.Fatal(err)
		}
		uview.UserList(response.Payload, userId)
	}()

	return <-userId
}
