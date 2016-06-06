package main

import (
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/kovetskiy/stash"
)

func createRepository(
	client stash.Stash, baseURL *url.URL,
	project, repository string,
) error {
	repo, err := client.CreateRepository(project, repository)
	if err != nil {
		return err
	}

	fmt.Printf(
		"%s/repos/projects/%s/repos/%s\n",
		strings.TrimRight(baseURL.String(), "/"),
		repo.Project.Key,
		repo.Name,
	)

	for _, link := range repo.Links.Clones {
		fmt.Printf("clone/%s: %s\n", link.Name, link.HREF)
	}

	return nil
}

func listRepositories(
	client stash.Stash, baseURL *url.URL,
	project string,
) error {
	repos, err := client.GetRepositories()
	if err != nil {
		return err
	}

	repositories := []string{}
	for _, repo := range repos {
		if strings.ToLower(repo.Project.Key) == strings.ToLower(project) {
			repositories = append(repositories, repo.Name)
		}
	}

	sort.Strings(repositories)

	for _, repository := range repositories {
		fmt.Printf("%s\n", repository)
	}

	return nil
}
