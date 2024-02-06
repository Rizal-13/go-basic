package main

import "fmt"

func random() interface{}{
	return false
}

func main(){

	// Cara 1 tetapi terbatas pada 1 type saja
	// result := random()
	// resultString := result.(string)
	// fmt.Println(result)
	// fmt.Println(resultString)

	// Cara 2 menggunakan switch (percabangan)

	result := random()
	switch value := result.(type){
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	// case bool:
	// 	fmt.Println("Bool", value)
	default:
		fmt.Println("Unknown")
	}


}