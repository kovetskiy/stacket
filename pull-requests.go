package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/kovetskiy/stash"
)

func handlePullRequestsCreate(
	args map[string]interface{},
) error {
	var (
		toProject         = args["<project>"].(string)
		toRepository      = args["<repository>"].(string)
		toBranch          = args["<branch>"].(string)
		fromProject, _    = args["<from-project>"].(string)
		fromRepository, _ = args["<from-repository>"].(string)
		fromBranch, _     = args["<from-branch>"].(string)
		title, _          = args["--title"].(string)
		description, _    = args["--desc"].(string)
		reviewers, _      = args["--reviewer"].([]string)
	)

	if fromProject == "" {
		fromProject = toProject
	}

	if fromRepository == "" {
		fromRepository = toRepository
	}

	if toBranch == "" {
		toBranch = "master"
	}

	if title == "" {
		title = fromBranch + " -> " + toBranch
	}

	fromRef := stash.PullRequestRef{
		Id: fromBranch,
		Repository: stash.PullRequestRepository{
			Slug: fromRepository,
			Project: stash.PullRequestProject{
				Key: fromProject,
			},
		},
	}
	toRef := stash.PullRequestRef{
		Id: toBranch,
		Repository: stash.PullRequestRepository{
			Slug: toRepository,
			Project: stash.PullRequestProject{
				Key: toProject,
			},
		},
	}

	pr, err := remote.CreatePullRequest(
		title, description,
		fromRef, toRef,
		reviewers,
	)
	if err != nil {
		return err
	}

	log.Printf(
		`The pull-request "%s -> %s" at %s/%s has been created:\n`,
		fromBranch, toBranch, toProject, toRepository,
	)

	fmt.Printf(
		"%s/projects/%s/repos/%s/pull-requests/%d/overview\n",
		strings.TrimRight(remote.url.String(), "/"),
		pr.ToRef.Repository.Project.Key, pr.ToRef.Repository.Slug,
		pr.ID,
	)

	return nil
}
