package goDemos

func PrintDemo2() {
	job := "r62"
	println(string(job[1]))
	job1 := string([]byte(job)[1:])
	println(job1)
}
