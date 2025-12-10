package main

import "fmt"

const UsdToRub float64 = 77.9
const UsdToEur float64 = 0.856

func main() {
	fmt.Printf("%f\n", UsdToRub)
	fmt.Printf("%f\n", UsdToEur)
	fmt.Println("EurToRub = ", UsdToRub/UsdToEur)
}
