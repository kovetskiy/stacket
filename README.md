stacket
=======

The client for Atlassiah Stash/Bitbucket Server that provides command line
interface for managing pull-requests, repositories and their settings.

## Features

- [x] Create a project
- [x] List repositories in a project
- [x] Create a repository in a project
- [x] Rename a repository
- [x] Move a repository from one project to another
- [x] Create a pull request
- [x] Uninstall an addon by key
- [x] Install an addon by given path
- [x] Disable an addon by key
- [x] Enable an addon by key
- [x] Set addon license

## Installation

**stacket** is go-gettable:

```
go get github.com/kovetskiy/stacket
```

Arch Linux users can create package using `pkgbuild` branch.

## Usage

```
Usage:
  stacket [options] projects       create    <project>
  stacket [options] repositories   list      <project>
  stacket [options] repositories   create    <project> <repository>
  stacket [options] repositories   rename    <project> <repository> <new-name>
  stacket [options] repositories   move      <project> <repository> <new-project>
  stacket [options] repositories   remove    <project> <repository>
  stacket [options] pull-requests  create    <project> <repository> <from> [<to>] [-r <reviewer>]...
  stacket [options] addons         uninstall <addon>
  stacket [options] addons         install   <path>
  stacket [options] addons         disable   <addon>
  stacket [options] addons         enable    <addon>
  stacket [options] addons license set       <addon>
  stacket -h |--help
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
                                 [default: /home/operator/.config/stacket.conf].
  --uri <bitbucket>             Use this URI instead of config.
  -g --git                      Set git remote origin.
  -h --help                     Show this screen.
  --version                     Show version.
```


## Configuration

You can write configuration file using following syntax:

```toml
base_url = "http://git.local"
username = "e.kovetskiy"
password = "sup3rp@ssw0rd31337"
```

or you are working with lots of bitbucket setups, you can pass the config as
URI to `--uri` flag:

```
stacket --uri http://admin:adminpass@bitbucket.local/
```

## Slack

Live chat in Slack: [slack.reconquest.io](https://slack.reconquest.io/)
