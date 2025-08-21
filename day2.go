package main

import (
	"GoLang/equation"
	"errors"
	"fmt"
)

func errorName(name string) error{
	if name == ""{
		return errors.New("Name must be filled")
	}
	return nil
}

func errorNum(age int) error{
	if age < 3{
		return fmt.Errorf("Age is too low : %d", age)
	}
	return nil
}

func noReturn() {
	result := satuReturn(5, 10)
	fmt.Println("Hello", result)
}

func satuReturn(a int, b int) (result int) {
	result = a + b
	return result
}

func multipleReturn(a int, b int) (hasil int, status string) {
	status = "SHAZAM!"
	hasil = a + b
	return hasil, status
}

func infiniteInput(num ...int) int {
	total := 0
	for _, value := range num {
		total += value
	}
	return total
}

func errorCheck(name string) error{
	if name == ""{
		return errors.New("Name must be filled")
	}
	return nil
}

func main() {
	defer func() {
		if r := recover(); r != nil{
			fmt.Println("Recovered from panic: ", r)
		}
	}() //NEEDS braces at the end buat immediately call the function, karena dia gada nama
	
	fmt.Println("Before")
	panic("AAAAA")


	err := errorCheck("")
	if err != nil {
		fmt.Println(err.Error())
	}

	// var err error = errorName("")
	// if err != nil{
	// 	fmt.Println(err.Error())
	// }
	// errorNum(1)
	noReturn()
	fmt.Println(satuReturn(2, 1))
	value, status := multipleReturn(2, 2)
	fmt.Printf("Value = %d, Status = %s\n", value, status)
	oneValue := infiniteInput(1, 2, 3, 4, 5)
	fmt.Println(oneValue)
	fmt.Println(equation.Kurang(10, 2))
	equation.Hello()
}
