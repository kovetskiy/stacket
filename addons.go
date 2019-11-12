package main

import (
	"fmt"

	"github.com/reconquest/karma-go"
)

func handleAddonsUninstall(
	args map[string]interface{},
) error {
	var name = args["<addon>"].(string)

	token, err := remote.GetUPMToken()
	if err != nil {
		return karma.Format(
			err,
			"unable to retrieve UPM token",
		)
	}

	err = remote.UninstallAddon(token, name)
	if err != nil {
		return karma.Format(
			err,
			"unable to uninstall addon",
		)
	}

	fmt.Printf("addon has been uninstalled: %s\n", name)

	return nil
}

func handleAddonsInstall(
	args map[string]interface{},
) error {
	var (
		path = args["<path>"].(string)
	)

	token, err := remote.GetUPMToken()
	if err != nil {
		return karma.Format(
			err,
			"unable to retrieve UPM token",
		)
	}

	err = remote.InstallAddon(token, path)
	if err != nil {
		return karma.Format(
			err,
			"unable to install addon",
		)
	}

	fmt.Printf("addon has been installed: %s\n", path)

	return nil
}
