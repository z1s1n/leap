package config

import (
	"io/ioutil"

	"github.com/chazyu1996/leap/tools"
	"gopkg.in/yaml.v2"
)

type Map map[string]interface{}
type Config struct {
	HostList        []HostList        `yaml:"hostList"`
	KeyList         []KeyList         `yaml:"keyList"`
	KeyWithPassList []KeyWithPassList `yaml:"keyWithPasswordList"`
}

type Host struct {
	Hostname string `yaml:"hostname"`
	Address  string `yaml:"address"`
}

func GetYamlConfig(filePath string) *Config {
	content, _ := ioutil.ReadFile(filePath)
	config := Config{}
	err := yaml.Unmarshal(content, &config)
	if err != nil {
		tools.PrintRed(err)
	}
	return &config
}

func GetHostByConfig(config *Config) []Map {
	var result []Map
	result = append(result, GetHostConfig(config)...)
	result = append(result, GetKeyConfig(config)...)
	result = append(result, GetKeyWithPassConfig(config)...)
	return result
}
