package main

import "fmt"

func main() {
	/*
		//variable here
		//Manifest Typing
		//deklarasi value langsung di variable
		var FirstName string = "Ditta"
		var Age int = 25
		//deklarasi value tidak langsung di variable, melainkan dibawah
		var LastName string
		LastName = "Amelia"

		//Variable Inference
		address := "Jakarta"
	*/
	//Multi Variable with Inference Technique
	FirstName, LastName, Age, address := "Ditta", "Amelia", 25, "Jakarta"

	//numeric varaible + value
	value1 := 192829
	var exist bool = true

	//variable with grave accent
	var message = `
	hi, i'm ditta, the quick
	brown fox jump over the lazy dog
	kmzwa8awaa-asd-dd `

	//constanta variable
	const phi = 3.14

	//math operation example (with static value)
	val1 := 1
	val2 := 3

	//All output
	fmt.Println("Hello World,", FirstName, LastName, " Age : ", Age)
	fmt.Println(address)
	fmt.Printf("Value 1 : %d \n", value1)
	fmt.Printf("Exist? %t \n", exist)
	fmt.Println(message, phi)
	fmt.Println("result : ", val1+val2)

}
