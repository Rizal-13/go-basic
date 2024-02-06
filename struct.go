package main

import "fmt"

type Customer struct {
	Name, Addr string
	Age        int
}

// struct method
func (customer Customer) sayHello(name string){
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {

	var inzun Customer
	inzun.Name = "Master"
	inzun.Addr = "Surabaya"
	inzun.Age = 25
	fmt.Println(inzun)


	crypto := Customer{
		Name : "crypto",
		Addr: "Surabaya",
		Age: 20,
	}
	fmt.Println(crypto)

	insun := Customer{"insun", "Surabaya", 27}
	fmt.Println(insun)
	fmt.Println(insun.Name)
	fmt.Println(insun.Addr)
	fmt.Println(insun.Age)

	// sayHello("Insun") => tidak dapat dieksekusi

	insun.sayHello("rizal")


}