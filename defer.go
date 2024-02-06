package main

import "fmt"

func logging(){
	fmt.Println("Selesai Memanggil Function")
}

func runApp(){
	defer logging()
	fmt.Println("Application just Run")
}

func main(){

	runApp()

}