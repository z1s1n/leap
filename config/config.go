package config

import (
	"io/ioutil"

	"github.com/chazyu1996/leap/tools"
	"gopkg.in/yaml.v2"
)

type Config struct {
	HostList []hostList `yaml:"hostList"`
}
type hostList struct {
	Port           int              `yaml:"port"`
	Authentication []authentication `yaml:"authentication"`
}

type authentication struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     []host `yaml:"host"`
}
type host struct {
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

func GetHostByConfig(config *Config) []map[string]interface{} {
	var result = []map[string]interface{}{}
	for _, hostList := range config.HostList {
		port := hostList.Port
		for _, auth := range hostList.Authentication {
			username := auth.Username
			password := auth.Password
			for _, host := range auth.Host {
				hostname := host.Hostname
				address := host.Address
				hostDict := make(map[string]interface{})
				hostDict["port"] = port
				hostDict["username"] = username
				hostDict["password"] = password
				hostDict["hostname"] = hostname
				hostDict["address"] = address

				result = append(result, hostDict)
			}
		}
	}
	return result
}
