package main 

import "fmt"

func main() {
	var firstVar string = "First Var"
	secondVar := "Second One"
	fmt.Println(firstVar, secondVar)


	var (
		thirdVar string = "ThirdVar ,"
		fourthVar string = "FourthVar"
	)

	fmt.Println(thirdVar, fourthVar)

	var5, var6 := "Variavel 5", "Variavel 6"
	fmt.Println(var5, var6)

	const imutable string = "N pode mudar"
	imutable = "Eu"
	
	fmt.Println(imutable)
}