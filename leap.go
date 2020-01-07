package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/chazyu1996/leap/client"
	"github.com/chazyu1996/leap/config"
	"github.com/chazyu1996/leap/tools"
	"github.com/chazyu1996/leap/utils"
)

func main() {
	configFilePath := flag.String("c", "config.yaml", "config path")
	flag.Parse()
	conf := config.GetYamlConfig(*configFilePath)
	utils.InitConfigList(conf)
	num := getInput()
	num -= 1
	host := utils.GetHost(num)
	address := host["address"].(string)
	port := strconv.Itoa(host["port"].(int))
	username := host["username"].(string)
	password := host["password"].(string)
	cli, err := client.DialWithPasswd(address+":"+port, username, password)
	if err != nil {
		fmt.Println(err)
	}
	defer cli.Close()

	err = client.New(cli)
	if err != nil {
		fmt.Println(err)
	}
}

func getInput() int {
	utils.Nav()
	numStr := tools.GetInput()
	num, err := strconv.Atoi(numStr[:len(numStr)-1])
	if err != nil {
		utils.Search(numStr[:len(numStr)-1])
		num = getInput()
	}
	return num
}
