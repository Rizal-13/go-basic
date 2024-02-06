package main

import "fmt"

// variadic :variabel argument
func sumAll(numbers ... int) int{
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}


func main(){
	result := sumAll(1,2,3,4,5,6)
	fmt.Println(result)
	
	fmt.Println(sumAll(7,8,9,10))

	// apabila dalam sebuah variabel terdapat slice maka dapat ditambahkan ... untuk merubah nya ke variadic
	nomer := []int {1,2,3,4,5,6}
	fmt.Println(sumAll(nomer...))
}