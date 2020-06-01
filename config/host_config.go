package config

type HostList struct {
	Port int    `yaml:"port"`
	Auth []auth `yaml:"auth"`
}

type auth struct {
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	GoogleAuthToken string `yaml:"googleAuthToken,omitempty"`
	Host            []Host `yaml:"host"`
}

func GetHostConfig(config *Config) *[]AllConfig {
	var result []AllConfig
	for _, hostList := range config.HostList {
		for _, auth := range hostList.Auth {
			for _, host := range auth.Host {
				hconf := AllConfig{
					Port:            hostList.Port,
					Hostname:        host.Hostname,
					Username:        auth.Username,
					Password:        auth.Password,
					Address:         host.Address,
					GoogleAuthToken: auth.GoogleAuthToken,
					Type:            "host",
				}
				result = append(result, hconf)
			}
		}
	}
	return &result
}
