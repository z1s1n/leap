package config

type KeyWithPassList struct {
	Port int               `yaml:"port"`
	Auth []KeyWithPassAuth `yaml:"auth"`
}
type KeyWithPassAuth struct {
	Username string `yaml:"username"`
	KeyFile  string `yaml:"keyFile"`
	Password string `yaml:"password"`
	Host     []Host `yaml:"host"`
}

func GetKeyWithPassConfig(config *Config) []Map {
	var result []Map
	for _, hostList := range config.KeyWithPassList {
		port := hostList.Port
		for _, auth := range hostList.Auth {
			username := auth.Username
			password := auth.Password
			keyFile := auth.KeyFile
			for _, host := range auth.Host {
				hostname := host.Hostname
				address := host.Address
				hostDict := make(Map)
				hostDict["port"] = port
				hostDict["username"] = username
				hostDict["keyFile"] = keyFile
				hostDict["password"] = password
				hostDict["hostname"] = hostname
				hostDict["address"] = address
				hostDict["type"] = "keyWithPass"
				result = append(result, hostDict)
			}
		}
	}
	return result
}
