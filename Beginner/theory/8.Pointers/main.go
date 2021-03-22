package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
	fmt.Println(&xPtr, "a")
	y := 42
	*xPtr = y

}

func square(x *float64) {
	*x = *x * *x
}
func main() {
	x := 5
	zero(&x)
	fmt.Println(x)  // x is 0
	fmt.Println(&x) // x is 0

	// using var keyword
	// we are not defining
	// any type with variable
	var y = 458

	// taking a pointer variable using
	// var keyword without specifying
	// the type
	var p = &y

	fmt.Println("Value stored in y before changing = ", y)
	fmt.Println("Address of y = ", &y)
	fmt.Println("Value stored in pointer variable p = ", p)

	// this is dereferencing a pointer
	// using * operator before a pointer
	// variable to access the value stored
	// at the variable at which it is pointing
	fmt.Println("Value stored in y(*p) Before Changing = ", *p)

	// changing the value of y by assigning
	// the new value to the pointer
	*p = 500

	fmt.Println("Value stored in y(*p) after Changing = ", y)

	q := 1.5
	square(&q)
	fmt.Println(q)

	one := 1
	two := 2
	swap(&one, &two)
	fmt.Println("now one==", one, "\nand two==", two)
}
func swap(x, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}
