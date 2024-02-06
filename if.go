package main

import "fmt"


func main(){
// if dan else
name := "insun"

if name == "insun"{
	fmt.Println("Hello Insun")
} else{
	fmt.Println("Hello Master")
}

// if , else, dan else if
names := "insun"

if names == "insun"{
	fmt.Println("Hello Insun")
} else if names == "insun123"{
	fmt.Println("Hello Insun")
} else if names == "insunaja"{
	fmt.Println("Hello Insun")
} else{
	fmt.Println("Hello Master")
}

if length := (len(names)) ; length <= 5{
	fmt.Println("Nama Sudah Ok !")
} else{
	fmt.Println("Nama Terlalu Panjang")
}

}