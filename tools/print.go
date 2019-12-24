package tools

import (
	"fmt"
)

func PrintGreen(a ...interface{}) {
	fmt.Printf("\033[1;32m%s\033[0m\n", a...)
}

func PrintRed(a ...interface{}) {
	fmt.Printf("\033[1;31m%s\033[0m\n", a...)
}
func PrintBlue(a ...interface{}) {
	fmt.Printf("\033[1;36m%s\033[0m\n", a...)
}

func PrintYellow(a ...interface{}) {
	fmt.Printf("\033[1;33m%s\033[0m\n", a...)
}

func StringCovert(str string, num int, after bool) string {
	defaultValue := ""
	for i := 0; i < num; i++ {

		defaultValue += " "
	}
	if after {
		str = str + defaultValue
	} else {
		str = defaultValue + str
	}
	return str
}
