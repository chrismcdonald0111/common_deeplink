# Simple HTTP Server

## Build
```
go build
```

## Compile Tests (caches test binaries)
```
// Current directory)
go test -c .
```

## Run Tests
```
// Current & subdirectories)
go test ./...
```

## Check Test Coverage
```
go test -cover ./...
```

## Run Server
```
// Use default config file
./redirect_server

// Pass config file as arg
./redirect_server "config/config.prod.json"
```

## Deploy Script
```
./bin/start.sh "config/config.prod.json"
```

## Stop Script
```
./bin/stop.sh
```
