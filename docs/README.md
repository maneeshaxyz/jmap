# jmap
<!--[![Go Report Card](https://goreportcard.com/badge/github.com/maneeshaxyz/jmap)](https://goreportcard.com/report/github.com/maneeshaxyz/jmap)-->

JMAP (JSON Meta Application Protocol) is a modern, efficient protocol for synchronizing email, calendars, and contacts over HTTP using JSON. It handles large datasets efficiently, supports push updates, and provides a consistent, easy-to-use interface across platforms.

### Build & run (source or dockerfile)

- Clone the repository using 
```
git clone github.com/maneeshaxyz/jmap
```

- from source:

```
make
```

- as dockerfile

```
make drun
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
