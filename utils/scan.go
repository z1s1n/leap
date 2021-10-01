package utils

import (
	"bufio"
	"os"
	"strconv"

	"github.com/z1son/leap/config"
)

var (
	String string
	Number int
	Input  string
)

func Scan() string {
	f := bufio.NewReader(os.Stdin) //读取输入的内容
	for {
		Nav()
		Input, _ = f.ReadString('\n') //定义一行输入的内容分隔符。
		if len(Input) == 1 {
			InitConfigList(config.ConfigInit)
			continue //如果用户输入的是一个空行就让用户继续输入。
		}
		return Input
	}
}
func GetInput() int {
	numStr := Scan()
	num, err := strconv.Atoi(numStr[:len(numStr)-1])
	if err != nil {
		Search(numStr[:len(numStr)-1])
		num = GetInput()
	}
	if num > len(*configList) {
		Search(numStr[:len(numStr)-1])
		num = GetInput()
	}
	return num
}
