package goDemos

import "fmt"

func PrintDemo5() {
	fmt.Println([]rune("世界你好"))
	//string to rune
	for _, char := range []rune("世界你好") {
		fmt.Println(string(char))
	}
}
