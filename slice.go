package main

import "fmt"

func main(){
	names := [...] string {"insun", "crypto", "47", "master", "good", "taichi"}

	slice1 := names[3:5]
	fmt.Println(slice1)

	slice2 := names[:5]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	days := [...] string {"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1 [0] = "sabtuLama"
	daySlice1 [1] = "mingguLama"
	fmt.Println(daySlice1)
	fmt.Println(days)

	// menambahkan slice baru ke array
	daySlice2 := append(daySlice1, "Libur Nasional")
	fmt.Println(daySlice2)

	daySlice2 [0] = "sabtuBaru"
	daySlice2 [1] = "mingguBaru"
	fmt.Println(daySlice2)

	// membuat slice baru
	newSlice := make([]string, 2, 5)
	newSlice[0] = "insun"
	newSlice[1] = "insun"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// newSlice2 := append(newSlice, "master")
	newSlice2 := append(newSlice, "master", "good")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "crypto"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// perbedaan array dan slice terletak di []
	iniArray := [...]int{1,2,3,4,5,6}
	iniSlice := []int{1,2,3,4,5,6}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}