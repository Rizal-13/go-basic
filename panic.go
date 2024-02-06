package main

import "fmt"

func endApp(){
	fmt.Println("End Application")
}

func runApp(eror bool){
	defer endApp()
	if eror{
		panic("Your App Error")
	}
}

func main(){

	runApp(true)

}