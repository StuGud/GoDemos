package goDemos

import (
	"fmt"
	"strings"
)

func PrintDemo13() {
	content := "r1S2S3r4r5"
	temp1 := make([]string, 0)
	temp := strings.Split(content, "r")
	for _, v := range temp {
		if i := strings.Index(v, "S"); i >= 0 {
			v = v[:i]
		}
		if v != "" {
			temp1 = append(temp1, v)
		}

	}
	for _, v := range content {
		fmt.Println(v)
	}
	fmt.Println(temp1)
}
