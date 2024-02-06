package main

import (
	"belajar-golang/database"
	_ "belajar-golang/internal" // jika ingin menjalankan init saja tanapa ada argumen lain maka harus ditambahkan _ didepan import packagenya
	"fmt"
)


func main(){

	fmt.Println(database.GetDatabase())

}