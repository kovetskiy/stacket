package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/kovetskiy/executil"
	"github.com/kovetskiy/stash"
	"github.com/seletskiy/hierr"
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

	fmt.Println(repo.SshUrl())

	err = setRemote(repo.SshUrl())
	if err != nil {
		return hierr.Errorf(
			err, "the repository has been created, but can't setup git remote",
		)
	}

	return nil
}

func setRemote(url string) error {
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		return nil
	}

	remotes, _, err := executil.Run(
		exec.Command("git", "remote", "show", "-n", "origin"),
	)
	if err != nil {
		return err
	}

	if len(strings.TrimSpace(string(remotes))) != 0 {
		_, _, err = executil.Run(
			exec.Command("git", "remote", "set-url", "origin", url),
		)
		if err != nil {
			return err
		}
	} else {
		_, _, err = executil.Run(
			exec.Command("git", "remote", "add", "origin", url),
		)
		if err != nil {
			return err
		}
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
