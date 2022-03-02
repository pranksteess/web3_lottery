package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var (
	Cfg CfgServer
)

func InitCfg(configFilePath string) error {
	if configFilePath == "" {
		configFilePath = "./config/config.yaml"
	}
	yamlFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(yamlFile, &Cfg); err != nil {
		log.Fatal(err)
	}
	return nil
}

type CfgServer struct {
	WhaleETHAddr []string `yaml:"whale_eth_addr" json:"whale_eth_addr"`
	SearchMethod string   `yaml:"search_method" json:"search_method"`
}
