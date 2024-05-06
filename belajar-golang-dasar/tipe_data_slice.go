package main

import "fmt"

// Tipe data slice adalah potongan dari tipe data array
// Slice biasanya digunakan untuk menyimpan data berbentuk array
// dan mempermudah penambahan data ke dalam array
// Slice dan array selalu terkoneksi, dimana slice adalah data yang mengakses sebagian atau seluruh data di array

// Tipe Data Slice memiliki 3 data, yaitu pointer, length, capacity
// Pointer menunjuk ke data pertama di array
// Length menunjuk ke data terakhir di array
// Capacity menunjuk ke kapasitas array slice, dimana length tidak boleh lebih dari capacity

func main() {
	names := []string{"Tri", "Nugroho", "Yosef", "Irawan", "Dwi", "Wahyudi"}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	var slice4 []string = names[:]
	fmt.Println(slice4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:] // sabtu, minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Hari Besar")
	daySlice2[0] = "Sabtu Lagi"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Eko"
	newSlice[1] = "Kurniawan"
	// newSlice[2] = "Khannedy"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Wahyudi")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	// Copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// perbedaan slice dan array

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
