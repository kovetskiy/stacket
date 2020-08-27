package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/reconquest/executil-go"
	"github.com/reconquest/hierr-go"
)

func handleRepositoriesCreate(
	args map[string]interface{},
) error {
	var (
		project       = args["<project>"].(string)
		repository, _ = args["<repository>"].(string)
	)

	repo, err := remote.CreateRepository(project, repository)
	if err != nil {
		return err
	}

	fmt.Printf(
		"repository %s/%s has been created:\n",
		project, repository,
	)

	fmt.Printf(
		"%s/projects/%s/repos/%s\n",
		strings.TrimRight(remote.url.String(), "/"),
		repo.Project.Key,
		repo.Name,
	)

	fmt.Printf(
		"ssh: %s\n",
		repo.SshUrl(),
	)

	if args["--git"].(bool) {
		err = setRemote(repo.SshUrl())
		if err != nil {
			return hierr.Errorf(
				err, "can't setup git remote",
			)
		}
	}

	return nil
}

func handleRepositoriesList(
	args map[string]interface{},
) error {
	project := args["<project>"].(string)

	repos, err := remote.GetProjectRepositories(project)
	if err != nil {
		return err
	}

	repositories := []string{}
	for _, repo := range repos {
		repositories = append(repositories, repo.Name)
	}

	sort.Strings(repositories)

	for _, repository := range repositories {
		fmt.Println(repository)
	}

	return nil
}

func handleRepositoriesRename(args map[string]interface{}) error {
	var (
		project       = args["<project>"].(string)
		repository, _ = args["<repository>"].(string)
		newName       = args["<new-name>"].(string)
	)

	err := remote.RenameRepository(project, repository, newName)
	if err != nil {
		return err
	}

	fmt.Printf(
		"the repository %s/%s has been renamed to %s/%s\n",
		project, repository,
		project, newName,
	)

	return nil
}

func handleRepositoriesMove(args map[string]interface{}) error {
	var (
		project       = args["<project>"].(string)
		repository, _ = args["<repository>"].(string)
		newProject    = args["<new-project>"].(string)
	)

	err := remote.MoveRepository(project, repository, newProject)
	if err != nil {
		return err
	}

	fmt.Printf(
		"the repository %s/%s has been moved to %s/%s\n",
		project, repository,
		newProject, repository,
	)

	return nil
}

func handleRepositoriesRemove(args map[string]interface{}) error {
	var (
		project       = args["<project>"].(string)
		repository, _ = args["<repository>"].(string)
	)

	err := remote.RemoveRepository(project, repository)
	if err != nil {
		return err
	}

	fmt.Printf(
		"the repository %s/%s has been removed\n",
		project, repository,
	)

	return nil
}

func handleRepositoriesFork(args map[string]interface{}) error {
	var (
		project          = args["<project>"].(string)
		repository, _    = args["<repository>"].(string)
		newRepository, _ = args["<new-repository>"].(string)
	)

	fork, err := remote.ForkRepository(project, repository, newRepository)
	if err != nil {
		return err
	}

	log.Println("The repository has been forked")

	fmt.Printf(
		"%s/%s\n",
		fork.Project.Key, fork.Slug,
	)

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
