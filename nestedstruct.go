package main

import (
	"fmt"
)

type salary struct {
	basic, hra, ta float64
}

type employee struct {
	name, id string
	age      int
	wages    []salary
}

func main() {
	emp1 := employee{
		name: "Akanksha",
		id:   "emp01",
		age:  25,
		wages: []salary{
			salary{
				basic: 10000,
				hra:   2000,
				ta:    500
			},
			salary{
				basic: 15000,
				hra:   3000,
				ta:    1000,
			},
		},
	}
	fmt.Println(emp1.name)
	fmt.Println(emp1.id)
	fmt.Println(emp1.age)
	fmt.Println(emp1.wages[0])
	fmt.Println(emp1.wages[1])
}
