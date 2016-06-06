package main

import (
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"github.com/docopt/docopt-go"
	"github.com/kovetskiy/stash"
)

var (
	version = "1.0"
	usage   = os.ExpandEnv(`
stacket is the client for Atlassian Stash/Bitbucket Server.

You should write configuration file using following syntax:
    base_url = "http://git.local"
    username = "e.kovetskiy"
    password = "sup3rp@ssw0rd31337"

Usage:
    stacket [options] repositories  list   <project>
    stacket [options] repositories  create <project> <repository>
    stacket [options] pull-requests create <project> <repository> <from> [<to>] [-r <reviewer>]...
    stacket -h | --help
    stacket --version

Options:
    repositories            Work with <project> repositories.
    pull-requests           Work with <project>/<repository> pull-requests.
        -t <title>          Speicfy pull-request title.
        -d <description>    Specify pull-request description.
    --config <path>         Use specified config file
                             [default: $HOME/.config/stacket.conf].
    -h --help               Show this screen.
    --version               Show version.
`)
)

func main() {
	args, err := docopt.Parse(usage, nil, true, version, true, true)
	if err != nil {
		panic(err)
	}

	var (
		repositoriesMode  = args["repositories"].(bool)
		pullRequestsMode  = args["pull-requests"].(bool)
		listMode          = args["list"].(bool)
		createMode        = args["create"].(bool)
		projectSlug       = args["<project>"].(string)
		repositorySlug, _ = args["<repository>"].(string)
		fromBranch, _     = args["<from>"].(string)
		toBranch, _       = args["<to>"].(string)
		title, _          = args["-t"].(string)
		description, _    = args["-d"].(string)
		reviewers, _      = args["-r"].([]string)

		configPath = args["--config"].(string)
	)

	config, err := getConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		log.Fatal(err)
	}

	stash.Log = log.New(ioutil.Discard, "", log.Ldate|log.Ltime|log.Lshortfile)

	client := stash.NewClient(config.Username, config.Password, baseURL)

	switch {
	case repositoriesMode && listMode:
		err = listRepositories(client, baseURL, projectSlug)

	case repositoriesMode && createMode:
		err = createRepository(client, baseURL, projectSlug, repositorySlug)

	case pullRequestsMode && createMode:
		err = createPullRequest(
			client, baseURL,
			projectSlug, repositorySlug,
			fromBranch, toBranch,
			title, description, reviewers,
		)

	}

	if err != nil {
		log.Fatal(err)
	}
}
