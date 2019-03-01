lock_file="$GOPATH/src/rakuten-it.com/rakuten/redirect_server/redirect_server.lock"
pid=$(cat $lock_file)
kill -s TERM $pid
