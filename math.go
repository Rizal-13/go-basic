package main

import "fmt"

func main() {
	var a = 12
	var b = 23
	var c = a + b
	fmt.Println(c)

	var i = 10
	i+= 3
	fmt.Println(i)
	i+=2
	fmt.Println(i)
	i++
	fmt.Println(i)
}