package main

import "fmt"

func main() {

	//konversi nilai integer
	var nilai32 int32 = 32760
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)


	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// konversi byte ke string
	var name = "insunCrypto"
	var e = name[5]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}