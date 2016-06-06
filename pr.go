package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/kovetskiy/stash"
)

func createPullRequest(
	client stash.Stash, baseURL *url.URL,
	project, repository string,
	from, to string,
	title, description string,
	reviewers []string,
) error {
	if to == "" {
		to = "master"
	}

	if title == "" {
		title = from + " -> " + to
	}

	pr, err := client.CreatePullRequest(
		project, repository,
		title, description,
		from, to,
		reviewers,
	)
	if err != nil {
		return err
	}

	fmt.Printf(
		"%s/repos/%s/project/%s/pull-requests/%d/overview\n",
		strings.TrimRight(baseURL.String(), "/"),
		project, repository,
		pr.ID,
	)

	return nil
}
