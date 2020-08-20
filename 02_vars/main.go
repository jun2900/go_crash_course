package main

import "fmt"

func main() {
	//Using var
	//var name = "Jun"
	var age int32 = 37
	const isCool = true
	var size float32 = 2.3

	//Shorthand
	//name := "Jun"
	//email := "jeremiahmaramis00@gmail.com"

	name, email := "Jun", "jeremiahmaramis00@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)

}
