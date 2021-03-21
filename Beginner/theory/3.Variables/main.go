package main

import "fmt"

func main() {
	var x string = "Hello World"
	fmt.Println(x)

	fmt.Print("Enter a number: ")
	var input float64
	var input1 float64
	fmt.Scanln(&input, &input1)

	output := input * 2 * input1

	fmt.Println("output = ",output)
	xx := 5
	xx += 1
	fmt.Println("xx = ",xx)
	fmt.Print("Напишите температуру в Фаренгейтах: ")
	var f float64
	fmt.Scanln(&f)
	fmt.Print(toCelsius(f))
}

// из градусов Фаренгейта в градусы Цельсия. (C = (F - 32) * 5/9)
func toCelsius (f float64) float64 {
	c :=(f - 32) * 5/9
	return c
}