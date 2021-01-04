package utils

import (
	"regexp"

	"github.com/morningfish/leap/config"
	"github.com/morningfish/leap/tools"
)

func Search(search string) {
	var newConfigList []config.AllConfig
	for _, config_ := range *configList {
		matchHostname, err := regexp.MatchString("(.*)"+search+"(.*)", config_.Hostname)
		if err != nil {
			tools.PrintRed(err)
		}
		matchAddress, err := regexp.MatchString("(.*)"+search+"(.*)", config_.Address)
		if err != nil {
			tools.PrintRed(err)
		}
		if matchAddress || matchHostname {
			newConfigList = append(newConfigList, config_)
		}
	}
	configList = &newConfigList
}
