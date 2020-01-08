package utils

import (
	"regexp"
	"strconv"

	"github.com/chazyu1996/leap/tools"
)

func Search(search string) {
	var newConfigList = []map[string]interface{}{}
	for _, config := range configList {
		hostname := config["hostname"].(string)
		address := config["address"].(string)
		matchHostname, err := regexp.MatchString("(.*)"+search+"(.*)", hostname)
		if err != nil {
			tools.PrintRed(err)
		}
		matchAddress, err := regexp.MatchString("(.*)"+search+"(.*)", address)
		if err != nil {
			tools.PrintRed(err)
		}
		if matchAddress || matchHostname {
			newConfigList = append(newConfigList, config)
		}
	}
	configList = newConfigList
}
func GetInput() int {
	Nav()
	numStr := tools.GetInput()
	num, err := strconv.Atoi(numStr[:len(numStr)-1])
	if err != nil {
		Search(numStr[:len(numStr)-1])
		num = GetInput()
	}
	return num
}
