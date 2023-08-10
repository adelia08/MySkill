package main

//struct => properti, dapat menampung beberapa jenis tipe data

import "fmt"

type student struct {
	name  string
	grade int
}

func main() {
	var s1 student
	s1.name = "my skill"
	s1.grade = 2
	fmt.Println("name :", s1.name)
	fmt.Println("grade :", s1.grade)
}
