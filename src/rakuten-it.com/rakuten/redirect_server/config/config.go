package config

import (
  "encoding/json"
  "io/ioutil"
  "os"
  "path/filepath"

  sc "rakuten-it.com/rakuten/redirect_server/service"
)

type JsonConfig struct {
    Port    int        `json:"port"`
    Timeout sc.Timeout `json:"timeout"`
}

// Read Json formatted config file into struct 'JsonConfig'
func ReadJsonConfig(fileName string) (JsonConfig, error) {
  // Get canonical file path
  fullPath, err := filepath.Abs(fileName)
  if err != nil {
    return JsonConfig{}, err
  }

  // Read file
  f, err := os.Open(fullPath)
  if err != nil {
    return JsonConfig{}, err
  }
  defer f.Close()

  // Read file bytes into memory
  byteValue, _ := ioutil.ReadAll(f)

	var config JsonConfig
  // Convert the bytes into defined struct 'JsonConfig'
	json.Unmarshal([]byte(byteValue), &config)

  return config, f.Sync() // Sync returns error if file buffer remains open
}
