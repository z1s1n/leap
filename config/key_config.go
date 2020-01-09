package config

type KeyList struct {
	Port int       `yaml:"port"`
	Auth []KeyAuth `yaml:"auth"`
}

type KeyAuth struct {
	Username string `yaml:"username"`
	KeyFile  string `yaml:"keyFile"`
	Host     []Host `yaml:"host"`
}

func GetKeyConfig(config *Config) []Map {
	var result []Map
	for _, hostList := range config.KeyList {
		port := hostList.Port
		for _, auth := range hostList.Auth {
			username := auth.Username
			keyFile := auth.KeyFile
			for _, host := range auth.Host {
				hostname := host.Hostname
				address := host.Address
				hostDict := make(Map)
				hostDict["port"] = port
				hostDict["username"] = username
				hostDict["keyFile"] = keyFile
				hostDict["hostname"] = hostname
				hostDict["address"] = address
				hostDict["type"] = "key"
				result = append(result, hostDict)
			}
		}
	}
	return result
}
