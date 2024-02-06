package main

import (
	helper "belajar-golang/help"
	"fmt"
)

func main(){

	result := helper.SayHello("Insun")
	fmt.Println(result)
	fmt.Println(helper.Appplication)
	fmt.Println(helper.sayGoodbye("insun"))


}