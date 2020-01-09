package config

type HostList struct {
	Port int    `yaml:"port"`
	Auth []auth `yaml:"auth"`
}

type auth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     []Host `yaml:"host"`
}

func GetHostConfig(config *Config) []Map {
	var result []Map
	for _, hostList := range config.HostList {
		port := hostList.Port
		for _, auth := range hostList.Auth {
			username := auth.Username
			password := auth.Password
			for _, host := range auth.Host {
				hostname := host.Hostname
				address := host.Address
				hostDict := make(Map)
				hostDict["port"] = port
				hostDict["username"] = username
				hostDict["password"] = password
				hostDict["hostname"] = hostname
				hostDict["address"] = address
				hostDict["type"] = "host"
				result = append(result, hostDict)
			}
		}
	}
	return result
}
