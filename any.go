package main

import "fmt"

func Ups() interface{} {
	return 1
}

func Ops() any {
	return 2
}

func main() {

	kosong := Ups()
	fmt.Println(kosong)

	kosong2 := Ops()
	fmt.Println(kosong2)

}