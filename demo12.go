package goDemos

import (
	"fmt"
	"strings"
)

func PrintDemo12() {
	test1()
	content := "r3S22r67"
	temp := strings.FieldsFunc(content, split)
	fmt.Println(temp)
}

func split(s rune) bool {
	if s == 'r' || s == 'S' {
		return true
	}
	return false
}

func test1() {
	temp := strings.SplitAfter("a,b,c,d", ",")
	fmt.Println(temp) //[a, b, c, d]
}
