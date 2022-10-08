package main

import "fmt"

func main() {

	hello()
	addnumbers(3, 5, 7)

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
