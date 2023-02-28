## hello-world-file

A docker image for tinkering with printing files into HTML.

```
rounakdatta/hello-world-file
```

### Using using Docker

```
docker run --publish=6060:6060 --rm -e DIRECTORY="/tmp" rounakdatta/hello-world-file
```

```
GET http://localhost:6060
```

### Using directly

```
DIRECTORY="/tmp" go run main.go
```

```
GET http://localhost:6060
```

#### Deserializing JSONified files

There's an additional flag called `JSON_MODE` which when set to any value will attempt to deserialize the value present in each file. That way the HTML would present JSON files as beautified.