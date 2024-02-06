package main

import "fmt"

func main(){
	// cara 1
	var nama [3] string

	nama[0] = "insun"
	nama[1] = "crypto"
	nama[2] = "47"

	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])

	//cara2
	var name = [3] int{12, 23, 43}
	fmt.Println(name)
	fmt.Println(name[0])
	fmt.Println(name[1])

	//cara 3
	var names = [...] int{12, 23, 43, 45, 67}
	fmt.Println(names)

	// func
	var jumlah = len(names)
	fmt.Println(jumlah)
	fmt.Println(len(nama))
}