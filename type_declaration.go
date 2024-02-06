package main

import "fmt"

func main() {

	type noKtp string

	var ktpInsun noKtp = "12212313141"

	var contoh string = "45452244"
	var contohKtp noKtp = noKtp(contoh)

	fmt.Println(ktpInsun)
	fmt.Println(contohKtp)



}