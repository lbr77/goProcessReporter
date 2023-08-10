package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

//  TODO
// ReadConfig
//  WriteConfig

type Config struct {
	ApiURL     string   `yaml:"api_url"`
	ApiKey     string   `yaml:"api_key"`
	ReportTime int64    `yaml:"report_time"`
	Keywords   []string `yaml:"keywords"`
	Replace    []string `yaml:"replace"`
	ReplaceTo  []string `yaml:"replace_to"`
}

func ReadConfig(configPath string) Config {
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("Error reading config,", err)
		return Config{}
	}
	var config Config
	if err = yaml.Unmarshal(yamlFile, &config); err != nil {
		fmt.Println("Failed to parse file", err)
		return Config{}
	}
	return config
}