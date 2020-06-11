package goDemos

import "fmt"

func PrintDemo8() {
	slice := []int{1}
	slice = slice[1:]
	fmt.Println(len(slice))
}
