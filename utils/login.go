// 登录选项
package utils

import (
	"strconv"

	"github.com/chazyu1996/leap/config"
	"github.com/chazyu1996/leap/tools"
)

var configList []map[string]interface{}

func InitConfigList(conf *config.Config) {
	configList = config.GetHostByConfig(conf)
}

func Nav() {
	number := 0
	tools.PrintGreen("= num = | ========== hostname ========== | ===== username ===== | ===== address ======")
	for _, hostInfo := range configList {
		number += 1
		hostname, _ := hostInfo["hostname"].(string)
		username, _ := hostInfo["username"].(string)
		address, _ := hostInfo["address"].(string)
		num := strconv.Itoa(number)
		num = tools.StringCovert(num, 7-len(num), true)
		hostname = tools.StringCovert(hostname, 30-len(hostname), true)
		username = tools.StringCovert(username, 20-len(username), true)
		address = tools.StringCovert(address, 20-len(address), true)
		tools.PrintGreen(num + " | " + hostname + " | " + username + " | " + address)
	}
	tools.PrintGreen("")
	tools.PrintGreen("===============================  请输入机器的序号  ===============================")
	tools.PrintGreen("")
	tools.PrintGreen("")
}

func GetHost(num int) map[string]interface{} {
	return configList[num]
}
