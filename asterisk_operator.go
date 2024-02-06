package main

import "fmt"

type Address struct{
	City, Prov, Country string
}

func main(){

address1 := Address{"Surabaya", "East Java", "Indonesia"}
address2 := &address1

address2.City = "Gresik" // copy value tanpa merubah value awal (Pass by Value)

// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"} // asterisk operator mengganti value dari value pointer awal

fmt.Println(address1)
fmt.Println(address2)
}