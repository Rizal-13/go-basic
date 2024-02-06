package main

import "fmt"

func endApp(){
	fmt.Println("End Application")
	message := recover()
	fmt.Println("Terjadi Error", message)
}

func runApp(eror bool){
	defer endApp()
	if eror{
		panic("Your App Error")
	}
}

func main(){

	runApp(true)
	fmt.Println("Insun")

}