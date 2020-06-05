package goDemos

import "strings"

func PrintDemo3() {
	var strs []string
	strs = append(strs, "abc")
	strs = append(strs, "efg")
	strs = append(strs, "hij")
	println(strs)

	join := strings.Join(strs, "\n")
	println(join)
}
