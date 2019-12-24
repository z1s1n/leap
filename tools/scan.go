package tools

import (
	"bufio"
	"os"
)

var (
	String string
	Number int
	Input  string
)

func GetInput() string {
	f := bufio.NewReader(os.Stdin) //读取输入的内容
	for {
		Input, _ = f.ReadString('\n') //定义一行输入的内容分隔符。
		if len(Input) == 1 {
			continue //如果用户输入的是一个空行就让用户继续输入。
		}
		return Input
	}
}
