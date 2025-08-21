package main

import "fmt"

func main(){
	// var peta map[string]string = map[string]string{
	// 	"name" : "Sky",
	// 	"city" : "Jakarduh",
	// }

	person := map[string]string{
		"name" : "sky",
		"city" : "jakarta",
		"number" : "08129918",
	}

	delete(person, "number")

	fmt.Println(person)
}