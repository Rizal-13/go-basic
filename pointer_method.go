package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = man.Name + " san"
}

func main() {

	insun := Man{"insun"}
	insun.Married()

	fmt.Println(insun.Name)

}