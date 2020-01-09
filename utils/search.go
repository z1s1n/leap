package utils

import (
	"github.com/chazyu1996/leap/config"
	"github.com/chazyu1996/leap/tools"
	"regexp"
)

func Search(search string) {
	var newConfigList []config.Map
	for _, config_ := range configList {
		hostname := config_["hostname"].(string)
		address := config_["address"].(string)
		matchHostname, err := regexp.MatchString("(.*)"+search+"(.*)", hostname)
		if err != nil {
			tools.PrintRed(err)
		}
		matchAddress, err := regexp.MatchString("(.*)"+search+"(.*)", address)
		if err != nil {
			tools.PrintRed(err)
		}
		if matchAddress || matchHostname {
			newConfigList = append(newConfigList, config_)
		}
	}
	configList = newConfigList
}
