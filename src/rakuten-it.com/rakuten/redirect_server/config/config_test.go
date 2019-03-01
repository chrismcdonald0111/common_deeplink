package config

import (
  "encoding/json"
  "io/ioutil"
  "os"
  "testing"

  sc "rakuten-it.com/rakuten/redirect_server/service"
)

func TestReadJsonConfig(t *testing.T) {
  configFile := "config_test.json"
  jsonBlob   := []byte(`
    { "port": 7777,
      "timeout": {
        "readTimeoutSec":  5,
        "writeTimeoutSec": 10,
        "idleTimeoutSec":  60
      }
    }
  `)

  // Create matching 'JsonConfig' struct
  expectedTimeout := sc.Timeout{
    ReadTimeoutSec:  5,
    WriteTimeoutSec: 10,
    IdleTimeoutSec:  60,
  }

  expectedJsonConfig := JsonConfig{
    Port:    7777,
    Timeout: expectedTimeout,
  }

  // Create non-matching 'JsonConfig' struct
  unexpectedTimeout := sc.Timeout{
    ReadTimeoutSec:  7,
    WriteTimeoutSec: 10,
    IdleTimeoutSec:  70,
  }

  unexpectedJsonConfig := JsonConfig{
    Port:    8888,
    Timeout: unexpectedTimeout,
  }

  // Read jsonBlob into 'JsonConfig' struct
  jsonConfig := JsonConfig{}
  err := json.Unmarshal(jsonBlob, &jsonConfig)
  if err != nil {
      t.Errorf("Failed to unmarshall json with error: %s", err)
  }

  // Serialize jsonConfig and write to 'configFile'
  json, err := json.Marshal(jsonConfig)
  err = ioutil.WriteFile(configFile, json, 0644)

  // Defer remove file to end of function execution stack
  defer os.Remove(configFile)

  // Read newly created json config file
  config, err := ReadJsonConfig(configFile)
  if err != nil {
    t.Errorf("Failed to parse config file with error: %s", err)
  }

  // Matching case
  if config != expectedJsonConfig {
     t.Errorf("Return value was incorrect, got: %+v, want: %+v.", config, expectedJsonConfig)
  }

  // Non-matching case
  if config == unexpectedJsonConfig {
     t.Errorf("Return value was incorrect, the following values should not match: %+v, want: %+v.", config, unexpectedJsonConfig)
  }
}
