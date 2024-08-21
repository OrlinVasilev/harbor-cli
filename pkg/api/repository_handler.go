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

package api

import (
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/repository"
	"github.com/goharbor/harbor-cli/pkg/utils"
	log "github.com/sirupsen/logrus"
)

func RepoDelete(projectName, repoName string) error {
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return err
	}
	_, err = client.Repository.DeleteRepository(ctx, &repository.DeleteRepositoryParams{ProjectName: projectName, RepositoryName: repoName})

	if err != nil {
		return err
	}

	log.Infof("Repository %s/%s deleted successfully", projectName, repoName)
	return nil
}

func RepoInfo(projectName, repoName string) error {
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return err
	}

	response, err := client.Repository.GetRepository(ctx, &repository.GetRepositoryParams{ProjectName: projectName, RepositoryName: repoName})

	if err != nil {
		return err
	}

	utils.PrintPayloadInJSONFormat(response.Payload)
	return nil
}

func ListRepository(projectName string) (repository.ListRepositoriesOK, error) {
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return repository.ListRepositoriesOK{}, err
	}

	response, err := client.Repository.ListRepositories(ctx, &repository.ListRepositoriesParams{ProjectName: projectName})

	if err != nil {
		return repository.ListRepositoriesOK{}, err
	}

	return *response, nil

}
