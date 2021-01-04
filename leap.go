package main

import (
	"flag"
	"fmt"

	"github.com/morningfish/leap/client"
	"github.com/morningfish/leap/config"
	"github.com/morningfish/leap/utils"
)

func main() {
	configFilePath := flag.String("c", "config.yaml", "config path")
	flag.Parse()
	conf := config.GetYamlConfig(*configFilePath)
	utils.InitConfigList(conf)
	num := utils.GetInput()
	num -= 1
	host := utils.GetHost(num)
	cli, err := client.GetClient(host)
	if err != nil {
		fmt.Println(err)
	}
	defer cli.Close()

	err = client.New(cli)
	if err != nil {
		fmt.Println(err)
	}
}
