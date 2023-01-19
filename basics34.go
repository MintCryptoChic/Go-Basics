package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, root1, root2, dis, ima float64
	fmt.Println("Enter a, b, c of the quadratic equation :")
	fmt.Scanln(&a, &b, &c)
	dis = b*b - (4 * a * c)
	if dis > 0 {
		root1 = (-b + math.Sqrt(dis)/(2*a))
		root1 = (-b - math.Sqrt(dis)/(2*a))
		fmt.Println("Two distinct and real roots exist:", root1, "and", root2)
	} else if dis == 0 {
		root1 = (-b / (2 * a))
		root2 = (-b / (2 * a))
		fmt.Println("Two equal and real roots exist:", root1, "and", root2)
	} else if dis < 0 {
		root1 = -b / (2 * a)
		root2 = -b / (2 * a)
		ima = math.Sqrt(-dis) / 2 * a
		fmt.Println("Two distinct complex roots exist:", root1, "+", ima, "and", root2, "-", ima)
	}
}
