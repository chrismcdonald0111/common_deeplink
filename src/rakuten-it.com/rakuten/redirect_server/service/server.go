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
    maxReadTimeoutSec      = 15
    minReadTimeoutSec      = 1
    defaultWriteTimeoutSec = 10
    maxWriteTimeoutSec     = 30
    minWriteTimeoutSec     = 1
    defaultIdleTimeoutSec  = 60
    maxIdleTimeoutSec      = 180
    minIdleTimeoutSec      = 10
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

// Check if port number is within defined limits:
// between maxPortNum and minPortNum (inclusive)
func isValidPortNum(port int) bool {
  if(port >= minPortNum && port <= maxPortNum) {
    return true
  }
  return false
}

// Use defaultPortNum if port number is invalid
func getPortNum(port int) int {
  if(isValidPortNum(port)) {
    return port
  }
  return defaultPortNum
}

// Validate and get 'Timeout' struct
func getTimeout(t Timeout) Timeout {
  // Use default timeouts if Timeout struct is empty
  if (t == Timeout{}) {
    return Timeout{
      ReadTimeoutSec:  defaultReadTimeoutSec,
      WriteTimeoutSec: defaultWriteTimeoutSec,
      IdleTimeoutSec:  defaultIdleTimeoutSec,
    }

  } else {
    // Use defaultReadTimeoutSec if read timeout is outside of defined limits
    if (t.ReadTimeoutSec > maxReadTimeoutSec || t.ReadTimeoutSec < minReadTimeoutSec){
      t.ReadTimeoutSec = defaultReadTimeoutSec
    }
    // Use defaultWriteTimeoutSec if write timeout is outside of defined limits
    if (t.WriteTimeoutSec > maxWriteTimeoutSec || t.WriteTimeoutSec < minWriteTimeoutSec){
      t.WriteTimeoutSec = defaultWriteTimeoutSec
    }
    // Use defaultIdleTimeoutSec if idle timeout is outside of defined limits
    if (t.IdleTimeoutSec > maxIdleTimeoutSec || t.IdleTimeoutSec < minIdleTimeoutSec){
      t.IdleTimeoutSec = defaultIdleTimeoutSec
    }
  }
  return t
}

// Create a new instance of 'server' struct and
// return new 'server' struct
func Create(port int, t Timeout) server {
  // Create instance of NewServeMux to avoid global route conflicts
  // from DefaultServeMux
  serveMux := http.NewServeMux()

  p       := getPortNum(port)
  timeout := getTimeout(t)
  addr    := fmt.Sprintf("%s%d", ":", p)
  s       := &http.Server{
    Addr:         addr,                                  // Port format: ":8080"
    Handler:      serveMux,
    ReadTimeout:  timeout.ReadTimeoutSec * time.Second,  // TLS -> Req Body
    WriteTimeout: timeout.WriteTimeoutSec * time.Second, // TLS -> Response
    IdleTimeout:  timeout.IdleTimeoutSec * time.Second,  // HTTP keep-alive
  }

  return server{
      router:     serveMux,
      HttpServer: s,
      Port:       p,
      Timeout:    timeout,
  }
}
