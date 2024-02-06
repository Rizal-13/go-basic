package main

import "fmt"

func main() {

	var nilaiAkhir = 90
	var absen = 80

	var lulusNilai = nilaiAkhir > 80
	var lulusAbsen = absen >= 80

	var lulus = lulusNilai && lulusAbsen

	fmt.Println(lulus)

}