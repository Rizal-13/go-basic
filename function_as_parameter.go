package main

import "fmt"

// cara 1
// func sayHelloWithFilter(name string, filter func(string) string) {
// 	fmt.Println("Hello", filter(name))
// }

// cara 2
type Filters func(string) string

func sayHelloWithFilter(name string, filter Filters) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string{
	if name == "asu"{
		return "***"
	} else{
		return name
	}
}


func main(){

	sayHelloWithFilter("insun", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("asu", filter)

}