package main

import "fmt"

func NewMap(name string) map[string]string{
	if name == ""{
		return nil
	} else{
		return map[string]string{
			"name": name,
		}
	}
}

func main(){

	nama := NewMap("insun")
	fmt.Println(nama["name"])

	nama2 := NewMap("")
	if nama2 == nil {
		fmt.Println("Data Kosong")
	} else{
		fmt.Println(nama2["name"])
	}

}