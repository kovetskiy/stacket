stacket
=======

The client for Atlassiah Stash/Bitbucket Server that provides command line
interface for managing pull-requests, repositories and their settings.

## Features

- [x] repositories: create new
- [x] repositories: list
- [ ] repositories: edit settings
- [x] pull-requests: create new
- [ ] pull-requests: edit
- [ ] pull-requests: approve
- [ ] pull-requests: merge
- [ ] pull-requests: decline

## Installation

**stacket** is go-gettable:

```
go get github.com/kovetskiy/stacket
```

Arch Linux users can create package using `pkgbuild` branch.

## Usage

#### Create a repository

```
stacket repositories create <project> <repository>
```


#### List repositories

```
stacket repositories list <project>
```

#### Create a pull-request

```
stacket pull-requests create <project> <repository> \
    <from> [<to>] \
    [-t <title>] [-d <description>] \
    [-r <reviewer>]...
```

## Configuration

You should write configuration file using following syntax:

```toml
base_url = "http://git.local"
username = "e.kovetskiy"
password = "sup3rp@ssw0rd31337"
```
