package main

import "fmt"

func main(){
	var slice []string = []string{"Sky", "Tes", "Tis"}
	fmt.Println(slice[0:3])

	slice[0] = "Bob"
	fmt.Println(slice[0:3])

	testSlice := slice[0:3]
	testSlice[0] = "Gopal"
	testSlice[1] = "Aki"
	fmt.Println(testSlice[0:3])
	fmt.Println(slice[0:3])

	// newSlice := append(slice, "Testtt")
	// newSlice[0] = "Kera"
	// fmt.Println(newSlice[0:4])
	// fmt.Println(slice[0:3])

	newSlice := append(testSlice, "Testtt")
	newSlice[0] = "Kera"
	fmt.Println(newSlice[0:4])
	fmt.Println(testSlice[0:3])


	// newSlice[0] = "Goblin"
	// newSlice[1] = "Goper"
	// fmt.Println(newSlice[0:4])
	// fmt.Println(len(slice))
	// print(len(newSlice))
	// print(cap(slice))
	// print(cap(newSlice))

	var autoSlice []string = make([]string, 2, 5)
	autoSlice[0] = "Sky"
	autoSlice[1] = "Jason"

	makeSlice := make([]string, 2, 10)
	makeSlice[0] = "Tes1"
	makeSlice[1] = "Tes2"
	makeSlice = append(makeSlice, "Tes3")
	fmt.Println(makeSlice[0:3])
	fmt.Println(len(makeSlice))
	fmt.Println(cap(makeSlice))
}