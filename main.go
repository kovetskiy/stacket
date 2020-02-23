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
	version = "2.2"
	usage   = os.ExpandEnv(`
stacket is the client for Atlassian Stash/Bitbucket Server.

You should write configuration file using following syntax:
  base_url = "http://git.local"
  username = "e.kovetskiy"
  password = "sup3rp@ssw0rd31337"

Usage:
  stacket [options] projects      create <project>
  stacket [options] repositories  list   <project>
  stacket [options] repositories  create <project> <repository>
  stacket [options] repositories  rename <project> <repository> <new-name>
  stacket [options] repositories  move   <project> <repository> <new-project>
  stacket [options] repositories  remove <project> <repository>
  stacket [options] pull-requests create <project> <repository> <from> [<to>] [-r <reviewer>]...
  stacket [options] addons uninstall <addon>
  stacket [options] addons install <path>
  stacket [options] addons license set <addon>
  stacket -h | --help
  stacket --version

Options:
  projects                      Work with projects.
  repositories                  Work with <project> repositories.
  pull-requests                 Work with <project>/<repository> pull-requests.
    -t --title <title>          Speicfy pull-request title.
    -d --desc <description>     Specify pull-request description.
    -r --reviewer <reviewer>    Specify pull-request reviewer.
  addons                        Work with Atlassian Add-ons.
  --config <path>               Use specified config file
                                 [default: $HOME/.config/stacket.conf].
  --uri <bitbucket>             Use this URI instead of config.
  -g --git                      Set git remote origin.
  -h --help                     Show this screen.
  --version                     Show version.
`)
)

var remote = struct {
	stash.Stash
	url *url.URL
}{}

type handlers map[string]func(map[string]interface{}) error

func init() {
	stash.Log = log.New(ioutil.Discard, "", 0)
	log.SetFlags(0)
}

func main() {
	args, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		panic(err)
	}

	var config Config
	if rawuri, ok := args["--uri"].(string); ok {
		config, err = getConfigFromURI(rawuri)
	} else {
		config, err = getConfig(args["--config"].(string))
	}
	if err != nil {
		log.Fatal(err)
	}

	remote.url, err = url.Parse(config.BaseURL)
	if err != nil {
		log.Fatal(err)
	}

	remote.Stash = stash.NewClient(
		config.Username,
		config.Password,
		remote.url,
	)

	switch {
	case args["projects"]:
		err = handle(args, handlers{
			"create": handleProjectsCreate,
		})

	case args["repositories"]:
		err = handle(
			args,
			handlers{
				"list":   handleRepositoriesList,
				"create": handleRepositoriesCreate,
				"rename": handleRepositoriesRename,
				"move":   handleRepositoriesMove,
				"remove": handleRepositoriesRemove,
			},
		)

	case args["pull-requests"].(bool):
		err = handle(
			args,
			handlers{
				"create": handlePullRequestsCreate,
			},
		)

	case args["addons"].(bool):
		err = handle(args,
			handlers{
				"uninstall": handleAddonsUninstall,
				"install":   handleAddonsInstall,
				"license":   handleAddonsLicense,
			},
		)
	}

	if err != nil {
		log.Fatal(err)
	}
}

func handle(args map[string]interface{}, callbacks handlers) error {
	for arg, callback := range callbacks {
		if ok, _ := args[arg].(bool); ok {
			return callback(args)
		}
	}

	return nil
}
