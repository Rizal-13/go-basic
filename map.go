package main

import "fmt"


func main(){

	person := map[string] string {"name" : "insun", "addr" :"aja"}
	
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["addr"])

	book := map[string]string {"tittle" : "Sapiens", "author" : "Yuval Noah Harari", "tes" : "salah"}

	fmt.Println(book)

	delete(book, "tes")
	fmt.Println(book)
}