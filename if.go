package main

import "fmt"

func main(){
	
	name := "sky1234"

	// if name := "Sky"; name == "Sky"{
	// 	fmt.Println("True")
	// } else if name == "sky" {
	// 	fmt.Println("False")
	// } else {
	// 	fmt.Println("Miss")
	// }

	class := "B"

	switch class{
	case "A":
		fmt.Println("aaaaa")
	case "B":
		fmt.Println("bbbb")
	default:
		fmt.Println("yang bener aje")
	}

	switch length := len(name); length > 3{
	case true:
		fmt.Println("True")
	case false:
		fmt.Println("False")
	}
}