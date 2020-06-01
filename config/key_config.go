package config

type KeyList struct {
	Port int       `yaml:"port"`
	Auth []KeyAuth `yaml:"auth"`
}

type KeyAuth struct {
	Username        string `yaml:"username"`
	KeyFile         string `yaml:"keyFile"`
	GoogleAuthToken string `yaml:"googleAuthToken,omitempty"`
	Host            []Host `yaml:"host"`
}

func GetKeyConfig(config *Config) *[]AllConfig {
	var result []AllConfig
	for _, hostList := range config.KeyList {
		for _, auth := range hostList.Auth {
			for _, host := range auth.Host {
				hconf := AllConfig{
					Hostname:        host.Hostname,
					Port:            hostList.Port,
					Username:        auth.Username,
					Address:         host.Address,
					GoogleAuthToken: auth.GoogleAuthToken,
					KeyFile:         auth.KeyFile,
					Type:            "key",
				}
				result = append(result, hconf)
			}
		}
	}
	return &result
}
