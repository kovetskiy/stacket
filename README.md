stacket
=======

The client for Atlassiah Stash/Bitbucket Server.

## Features

- [x] repositories: create new
- [x] repositories: list
- [ ] repositories: edit settings
- [x] pull-requests: create new
- [ ] pull-requests: edit
- [ ] pull-requests: approve
- [ ] pull-requests: merge
- [ ] pull-requests: decline


## Usage

#### Create a repository
```
stacket repositories create <project> <repository>
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

