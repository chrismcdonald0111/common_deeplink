/*
  Author: Chris McDonald
  This project a simple web server written in Go. This project is supposed
  to server as a solid base/boilerplate for more complex web server
  implementations in Go. The main feature of this project is URL redirection.
  This project includes the following features:
  - Run-time argument parsing
  - JSON config <> struct parsing
  - STDOUT Loggingd
  - Framework-less web server (built on net/http)
  - HTTP protocol (HTTPS is not yet implemented)
  - URL Redirection

  TODO:
  - Add Godocs
  - Add support for HTTPS
*/

package main

import (
    "fmt"
    "log"
    "os"

    co "rakuten-it.com/rakuten/redirect_server/config"
    sv "rakuten-it.com/rakuten/redirect_server/service"
)

const (
  configFilePos = 0 // if ./cmd a b c then pos 1 would be a
  defaultConfigFile = "config/config.stg.json"
)

func parseArg(args []string, pos int, defaultArg string) string {
  if len(args) > pos {
    return args[pos]
  }
  return defaultArg
}

// Parse arguments & config
// Create instance of HTTP Server (server struct)
// Add routes to server mux
// Server begins listening on provided port
func main() {
    args := os.Args[1:] // Start from pos 1 to ignore the command literal i.e. ./cmd a b c would ignore ./cmd
    configFile := parseArg(args, configFilePos, defaultConfigFile)
    config, err := co.ReadConfig(configFile)
    if err != nil {
      log.Fatal(err)
    }
    s := sv.Create(config.Port, config.Timeout)
    fmt.Println(s.Port)
    s.AddRoutes()
    log.Fatal(s.HttpServer.ListenAndServe())
}
