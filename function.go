package main

import "fmt"

// membuat function
func sayHello(){
	fmt.Println("Hello World !")
}

// membuat function dan parameternya
func sayHelloTo(firstName string, lastName string){
	fmt.Println("Hello", firstName, lastName)
}

// membuat return value
func getHello(name string) string{
	hello := "Hello " + name
	return hello
}

// membuat return multiple
func getFullName() (string, string){
	return "insun", "crypto"
}

// membuat return name value dan dapat dikosongi value nya
func getComplete() (firstName, middleName, lastName string){
	firstName = "insun"
	middleName = "crypto"
	lastName = "47"
	return firstName, middleName, lastName
}

func getComplete1() (firstName string, middleName string, numer int){
	firstName = "insun"
	middleName = "crypto"
	numer = 47
	return firstName, middleName, numer
}

func main(){
	sayHello()
	sayHelloTo("insun", "crypto")

	//return value 1
	result := getHello("insun")
	fmt.Println(result)
	//return value 2
	fmt.Println(getHello("Master"))

	//return value multiple
	// first, last := getFullName()
	// fmt.Println(first, last)

	//return salah satu value multiple 
	_, last := getFullName()
	fmt.Println(last)

	a, b, c := getComplete()
	fmt.Println(a, b, c)

	d, e, f := getComplete1()
	fmt.Println(d, e, f)
}