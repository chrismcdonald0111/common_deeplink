cd "$GOPATH/src/rakuten-it.com/rakuten/redirect_server"
go build
go test ./...
nohup ./redirect_server $@ > redirect_server.out 2>&1 &
echo $! > redirect_server.lock
