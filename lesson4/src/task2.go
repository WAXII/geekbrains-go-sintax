package main

import (
	"fmt"
	"sort"
	"task2"
)

func main() {
	Book := task2.Book{
		&task2.ContactItem{
			Name:    "n1",
			SurName: "sn1",
			Phone:   82223334455,
		},
		&task2.ContactItem{
			Name:    "n2",
			SurName: "sn2",
			Phone:   85556667788,
		},
		&task2.ContactItem{
			Name:    "n3",
			SurName: "sn4",
			Phone:   80001112233,
		},
		&task2.ContactItem{
			Name:    "n4",
			SurName: "sn4",
			Phone:   83334445566,
		},
	}
	sort.Sort(Book)
	for _, c := range Book {
		fmt.Println(c.Name, c.SurName, ">", c.Phone)
	}
}
