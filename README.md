# jmap
<!--[![Go Report Card](https://goreportcard.com/badge/github.com/maneeshaxyz/jmap)](https://goreportcard.com/report/github.com/maneeshaxyz/jmap)-->

JMAP (JSON Meta Application Protocol) is a modern, efficient protocol for synchronizing email, calendars, and contacts over HTTP using JSON. It handles large datasets efficiently, supports push updates, and provides a consistent, easy-to-use interface across platforms.


### Setup
- Clone the repository using 
```
git clone github.com/maneeshaxyz/jmap
```

- generate a .cert and .key for testing using 
``` 
openssl req -x509 -newkey rsa:4096 -nodes -keyout server.key -out server.crt -days 365 -subj "/CN=localhost"
```

### Build (source or dockerfile)
- build from source:

```
make
```

- as dockerfile

```
make drun
```
