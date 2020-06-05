package goDemos

import "fmt"

//比较字符串相等

func PrintDemo7() {
	str1 := "abc"
	str2 := []string{"abc"}
	if str1 == str2[0] {
		fmt.Println("1字符串相等")
	}

}
