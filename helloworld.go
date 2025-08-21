package main

import "fmt"

func sayHello(name string, age int){
	fmt.Println("Hello "+ name + ", your age is:", age)
}

func hellopt2(age int) int{
	return age
}

func hellopt3(name string) string{
	return "Hello" + name
}

func main(){
	var array []string = []string{"Sky", "Tes", "Tis"}
	fmt.Println(array[0:3])
	

	sayHello("Sky", 20)
}