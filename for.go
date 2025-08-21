package main

import "fmt"

func main(){
	for i := 1; i < 10; i++{
		fmt.Println("Counter +", i)
	}

	var data []string = []string{"Tes1", "Tes2", "tes3"}
	// for i := 0; i<len(data); i++{
	// 	fmt.Println(data[i])
	// }

	for index, name := range(data){
		fmt.Println("Index", index+1, "Item", name)
	}
}