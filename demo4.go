package goDemos

import "strings"

func PrintDemo4() {
	job := "S6"
	if strings.HasPrefix(job, "S") {
		println("succeed")
	}
}
