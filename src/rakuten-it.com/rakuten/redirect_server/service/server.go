package service

import (
    "fmt"
    "net/http"
    "time"
)

const (
    defaultPortNum  = 8080
    maxPortNum      = 65535
    minPortNum      = 1025

    defaultReadTimeoutSec  = 5
    defaultWriteTimeoutSec = 10
    defaultIdleTimeoutSec  = 60
)

type server struct {
  router *http.ServeMux
  HttpServer *http.Server
  Port int
  Timeout Timeout
}

type Timeout struct {
  ReadTimeoutSec time.Duration
  WriteTimeoutSec time.Duration
  IdleTimeoutSec time.Duration
}

func isValidPortNum(port int) bool {
  if(port >= minPortNum && port <= maxPortNum) {
    return true
  }
  return false
}

func getPortNum(port int) int {
  if(isValidPortNum(port)) {
    return port
  }
  return defaultPortNum
}

func getTimeout(t Timeout) Timeout {
  if (t == Timeout{}) {
    return Timeout{
      ReadTimeoutSec:  defaultReadTimeoutSec,
      WriteTimeoutSec: defaultWriteTimeoutSec,
      IdleTimeoutSec:  defaultIdleTimeoutSec,
    }
  }
  return t
}

func Create(port int, t Timeout) server {
  serveMux := http.NewServeMux() // Create NewServeMux to avoid global route conflicts from DefaultServeMux

  p       := getPortNum(port)
  timeout := getTimeout(t)
  addr    := fmt.Sprintf("%s%d", ":", p)
  s       := &http.Server{
      Addr:         addr,
      Handler:      serveMux,
      ReadTimeout:  timeout.ReadTimeoutSec * time.Second,
      WriteTimeout: timeout.WriteTimeoutSec * time.Second,
      IdleTimeout:  timeout.IdleTimeoutSec * time.Second,
  }

  return server{
      router:     serveMux,
      HttpServer: s,
      Port:       p,
      Timeout:    timeout,
  }
}
