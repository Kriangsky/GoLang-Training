package main

import "fmt"

func main() {


	var(
		name string = "Sky"
		age int= 12
	)
	fmt.Println(name+" and", age)
	

	const pi = 3.14

	// Input
	var namaLengkap string
	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&namaLengkap)
	fmt.Println("Input kamu", namaLengkap)

	var umur int
	fmt.Print("Masukkan umur: ")
	fmt.Scanln(&umur)
	fmt.Println("Input kamu", umur)

	// Printf (Cannot use Println lor)
	fmt.Printf("Nama kamu %s dan umur kamu %d", namaLengkap, umur)
	fmt.Println()
	var names [5]string
	names[0] = "Sky"
	names[1] = "Tes"
	fmt.Println(names[:5])
	fmt.Println(len(names))

	var fruits []string = []string{"Sky1", "Sky2"}
	fmt.Println(fruits[0:2])
	fmt.Println(len(fruits))

	fruits = append(fruits, "Sky3")
	fmt.Println(fruits[0:3])
	fmt.Println(len(fruits))
	fruits = append(fruits, "Sky4")
	fruits = append(fruits, "Sky5") 
	fmt.Println(fruits[0:5])
	fmt.Println(len(fruits))

	items := make([]string, 3, 5)
	items[0] = "Sapi"
	items[1] = "Kambing"
	items[2] = "Ular"
	items = append(items, "Ayam")
	items = append(items, "Itik")
	items = append(items, "Ayam2")

	for _, items := range(items){
		fmt.Println(items)
	}

	fmt.Println(len(items), cap(items))

	// Loop
	for i := 1; i<5; i++{
		fmt.Println(i)
	}

	// Switch
	var hari string = "Minggu"

	switch hari{
	case "Senin":
		fmt.Println("1")
	case "Selasa":
		fmt.Println("2")
	case "Rabu":
		fmt.Println("3")
	case "Sabtu", "Minggu": //bisa 2 biji gini
		fmt.Println("WEEKENDDSSSSS")
	default:
		fmt.Println("Gak ada")
	}

	// Pointer
	var house *string
	var fill string = "Chair"
	house = &fill
	fmt.Println(*house)
	fmt.Println(house)

	// Struct
	type House struct{
		paint string
		houseAge int
		typesOfHouses []string
	}

	type Customer struct{
		name string
		age int
		itemsBought []string
		house House
	}

	var data Customer
	data.name = "Sky"
	data.age = 18
	data.itemsBought = []string{"苹果", "Apple", "Water"}
	data.house.paint = "Red"
	data.house.houseAge = 20
	data.house.typesOfHouses = []string{"Elevated", "Frank Lloyd Wright"}

	fmt.Println(data.house.typesOfHouses)
	// fmt.Println(data)
}
