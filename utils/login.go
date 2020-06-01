// 登录选项
package utils

import (
	"strconv"

	"github.com/chazyu1996/leap/config"
	"github.com/chazyu1996/leap/tools"
)

var configList *[]config.AllConfig

func InitConfigList(conf *config.Config) {
	configList = config.GetHostByConfig(conf)
}

func Nav() {
	number := 0
	tools.PrintGreen("= num = | ========== hostname ========== | ===== username ===== | ===== address ======")
	for _, hostInfo := range *configList {
		number += 1
		num := strconv.Itoa(number)
		num = tools.StringCovert(num, 7-len(num), true)
		hostname := tools.StringCovert(hostInfo.Hostname, 30-len(hostInfo.Hostname), true)
		username := tools.StringCovert(hostInfo.Username, 20-len(hostInfo.Username), true)
		address := tools.StringCovert(hostInfo.Address, 20-len(hostInfo.Address), true)
		tools.PrintGreen(num + " | " + hostname + " | " + username + " | " + address)
	}
	tools.PrintGreen("")
	tools.PrintGreen("===============================  请输入机器的序号  ===============================")
	tools.PrintGreen("")
	tools.PrintGreen("")
}

func GetHost(num int) config.AllConfig {
	return (*configList)[num]
}
