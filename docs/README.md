# jmap
<!--[![Go Report Card](https://goreportcard.com/badge/github.com/maneeshaxyz/jmap)](https://goreportcard.com/report/github.com/maneeshaxyz/jmap)-->

JMAP (JSON Meta Application Protocol) is a modern, efficient protocol for synchronizing email, calendars, and contacts over HTTP using JSON. It handles large datasets efficiently, supports push updates, and provides a consistent, easy-to-use interface across platforms.

### Setup
- Clone the repository using 
```
git clone github.com/maneeshaxyz/jmap
```

### Build & run (source or dockerfile)
- from source:

```
make #make build for build only
```

- as dockerfile

```
make drun # make dbuild for docker build only
```

## Project Structure

```
jmap/
├── cmd/jmapd/          # Main server entry point
├── internal/
│   ├── core/           # JMAP domain types
│   ├── parser/         # Request parsing and validation
│   └── server/         # HTTP server implementation
├── docs/               # All documentation
└── bin/                # Built binaries
```
