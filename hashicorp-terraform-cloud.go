package main

import (
	"context"
	"fmt"
	"net/url"
	"path"

	tfe "github.com/hashicorp/go-tfe"
)

func (config *hashicorpTerraformCloud) checkStatus() ([]result, error) {
	var results []result

	tfeURLBase, err := url.Parse("https://app.terraform.io/app")
	tfeURLBase.Path = path.Join(tfeURLBase.Path, config.OrgName)

	client, err := tfe.NewClient(&tfe.Config{
		Token: config.Token,
	})
	if err != nil {
		return results, err
	}

	ctx := context.Background()

	for _, w := range config.Workspaces {
		workspace, err := client.Workspaces.Read(ctx, config.OrgName, w)
		if err != nil {
			return results, err
		}

		lastRun, err := client.Runs.Read(ctx, workspace.CurrentRun.ID)
		if err != nil {
			return results, err
		}
		runURL := path.Join(tfeURLBase.String(), "workspaces", workspace.Name, "runs", lastRun.ID)

		if lastRun.Status == tfe.RunErrored {
			results = append(results, result{
				Name:        fmt.Sprintf("Hashicorp Terraform Cloud - %s", workspace.Name),
				Description: fmt.Sprintf("Last run %s at %s", lastRun.Status, lastRun.CreatedAt),
				Action:      runURL,
			})
		}
	}

	return results, nil
}
