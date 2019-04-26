package main

import (
	"fmt"
	"strings"
)

func handleProjectsCreate(
	args map[string]interface{},
) error {
	var (
		key = args["<project>"].(string)
	)

	project, err := remote.CreateProject(key)
	if err != nil {
		return err
	}

	fmt.Printf(
		"project %s has been created:\n",
		key,
	)

	fmt.Printf(
		"%s/repos/projects/%s\n",
		strings.TrimRight(remote.url.String(), "/"),
		project.Key,
	)

	return nil
}
