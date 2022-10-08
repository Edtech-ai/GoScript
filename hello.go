package main

import "fmt"

func main() {

	hello()
	addnumbers(3, 5, 7)
	result := double(10)
	fmt.Println(result)

}

func hello() {

	fmt.Println("Hello World!")
}

func addnumbers(x, y, z int) {

	a := x + y
	b := y + z
	c := x + z

	fmt.Println("The results are:", a, b, c)

}

func double(x int) int {

	y := x * 2
	return y
}
