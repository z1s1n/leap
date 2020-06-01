package config

type KeyWithPassList struct {
	Port int               `yaml:"port"`
	Auth []KeyWithPassAuth `yaml:"auth"`
}
type KeyWithPassAuth struct {
	Username        string `yaml:"username"`
	KeyFile         string `yaml:"keyFile"`
	Password        string `yaml:"password"`
	GoogleAuthToken string `yaml:"googleAuthToken,omitempty"`
	Host            []Host `yaml:"host"`
}

func GetKeyWithPassConfig(config *Config) *[]AllConfig {
	var result []AllConfig
	for _, hostList := range config.KeyWithPassList {
		for _, auth := range hostList.Auth {
			for _, host := range auth.Host {
				hconf := AllConfig{
					Hostname:        host.Hostname,
					Port:            hostList.Port,
					Username:        auth.Username,
					Address:         host.Address,
					GoogleAuthToken: auth.GoogleAuthToken,
					KeyFile:         auth.KeyFile,
					Password:        auth.Password,
					Type:            "keyWithPass",
				}
				result = append(result, hconf)
			}
		}
	}
	return &result
}
