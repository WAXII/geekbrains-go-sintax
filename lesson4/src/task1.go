package main

import "task1"

func main() {
	s1 := task1.Struct1{}
	s2 := task1.Struct2{}
	task1.TestStructToInterface(&s1)
	task1.TestStructToInterface(&s2)
}
