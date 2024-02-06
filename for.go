package main

import "fmt"


func main(){
	// cara 1
	count := 1

	for count <= 10{
		fmt.Println("Perulangan ke", count)
		count++
	}

	// cara 2
	for count2:=1; count2 <= 10; count2++{
		fmt.Println("Perulangan ke", count2)
		}

	// count3 := 15

	// for count3 >= 10{
	// 	fmt.Println("Perulangan ke", count3)
	// 	count3--
	// }

	// ini slice
	names := []string {"insun", "aja", "crypto"} 
	
	// range cara 1
	for i := 0; i< len(names); i++{
		fmt.Println(names[i])
	}

	// range cara 2
	for index, name := range names{
		fmt.Println("Index ke", index, "=", name)
	}

	// range cara 3 tanpa menggunakan index
	for _, name := range names{
		fmt.Println(name)
	}

}