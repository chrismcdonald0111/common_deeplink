package config

import (
  "encoding/json"
  "io/ioutil"
  "os"
  "path/filepath"

  sv "rakuten-it.com/rakuten/redirect_server/service"
)

type Config struct {
    Port    int        `json:"port"`
    Timeout sv.Timeout `json:"timeout"`
}

func ReadConfig(fileName string) (Config, error) {
  fullPath, _ := filepath.Abs(fileName)
  f, err := os.Open(fullPath)
  if err != nil {
    return Config{}, err
  }
  defer f.Close()

  byteValue, _ := ioutil.ReadAll(f)

	var config Config
	json.Unmarshal([]byte(byteValue), &config)

  return config, f.Sync()
}
