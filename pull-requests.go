package main

import (
	"fmt"
	"strings"
)

func handlePullRequestsCreate(
	args map[string]interface{},
) error {
	var (
		project        = args["<project>"].(string)
		repository, _  = args["<repository>"].(string)
		branchFrom, _  = args["<from>"].(string)
		branchTo, _    = args["<to>"].(string)
		title, _       = args["--title"].(string)
		description, _ = args["--desc"].(string)
		reviewers, _   = args["--reviewer"].([]string)
	)

	if branchTo == "" {
		branchTo = "master"
	}

	if title == "" {
		title = branchFrom + " -> " + branchTo
	}

	pr, err := remote.CreatePullRequest(
		project, repository,
		title, description,
		branchFrom, branchTo,
		reviewers,
	)
	if err != nil {
		return err
	}

	fmt.Printf(
		"%s/projects/%s/repos/%s/pull-requests/%d/overview\n",
		strings.TrimRight(remote.url.String(), "/"),
		project, repository,
		pr.ID,
	)

	return nil
}
