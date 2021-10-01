package config

import (
	"io/ioutil"

	"github.com/z1son/leap/tools"
	"gopkg.in/yaml.v2"
)

type Config struct {
	HostList        []HostList        `yaml:"hostList"`
	KeyList         []KeyList         `yaml:"keyList"`
	KeyWithPassList []KeyWithPassList `yaml:"keyWithPasswordList"`
}

type AllConfig struct {
	Username        string
	Password        string
	Hostname        string
	Address         string
	Port            int
	KeyFile         string
	GoogleAuthToken string
	Type            string
}
type Host struct {
	Hostname string `yaml:"hostname"`
	Address  string `yaml:"address"`
}

var ConfigInit *Config

func GetYamlConfig(filePath string) *Config {
	content, _ := ioutil.ReadFile(filePath)
	config := Config{}
	err := yaml.Unmarshal(content, &config)
	if err != nil {
		tools.PrintRed(err)
	}
	ConfigInit = &config
	return &config
}

func GetHostByConfig(config *Config) *[]AllConfig {
	var result []AllConfig
	result = append(result, *GetHostConfig(config)...)
	result = append(result, *GetKeyConfig(config)...)
	result = append(result, *GetKeyWithPassConfig(config)...)
	return &result
}
