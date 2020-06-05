package goDemos

import (
	"fmt"
	"strconv"
)

func PrintDemo6() {
	var charStr string
	var charRune rune
	var charInt int

	//string转rune
	fmt.Println("string转rune")
	charStr = "s"
	for i, r := range charStr {
		fmt.Println(i, r)
	}
	charRune = rune(charStr[0])
	fmt.Println(charRune)
	charInt = int(charStr[0])
	fmt.Println(charInt)

	//rune转int
	fmt.Println("rune转int")
	charRune = 's'
	fmt.Println(charRune)
	fmt.Println(int(charRune))

	//string转int
	fmt.Println("string转int")
	charStr = "ø"
	fmt.Println(int(charStr[0]))

	//int转rune
	fmt.Println("int转rune")
	charInt = 98
	fmt.Println(rune(charInt))
	fmt.Println(string(rune(charInt)))
	fmt.Println(string(charInt))

	//int转string
	charInt = 115
	fmt.Println("int转string")
	fmt.Println(strconv.Itoa(charInt))
	fmt.Println(string(charInt))

	//for i := 0;i <= 256; i++{
	//	fmt.Println(strconv.Itoa(i))
	//	fmt.Println(string(i)+"对应: "+string(rune(i)))
	//}

}
