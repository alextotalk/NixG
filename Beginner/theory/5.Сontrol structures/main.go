package main

import (
	"fmt"
)

func main() {
	OperatorFor()

	SwitchType()

	IfElse()

	SwitchOperator()

	EvenOdd()

	SelectOperator()

	FizzBuzz()
}
func OperatorFor() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println("____________")

}
func SwitchType() {
	var x interface{}

	//x = os.Args[1]
	//	x="say"
	x = 22

	switch i := x.(type) {
	case nil:
		fmt.Printf("type of x: %T\n", i)
	case int:
		fmt.Println("x is int")
	case float64:
		fmt.Println("x is float64")
	case func(int) float64:
		fmt.Printf("x is func(int)")
	case bool, string:
		fmt.Println("x is bool or string")
	default:
		fmt.Println("don't know the type")
	}
	fmt.Println("____________")
}
func IfElse() {
	/* local variable definition */
	var a int = 100

	/* check the boolean condition */
	if a == 10 {
		/* if condition is true then print the following */
		fmt.Printf("Value of a is 10\n")
	} else if a == 20 {
		/* if else if condition is true */
		fmt.Printf("Value of a is 20\n")
	} else if a == 30 {
		/* if else if condition is true  */
		fmt.Printf("Value of a is 30\n")
	} else {
		/* if none of the conditions is true */
		fmt.Printf("None of the values is matching\n")
	}
	fmt.Printf("Exact value of a is: %d\n", a)
	fmt.Println("____________")

}
func SwitchOperator() {

	var grade string
	// assign first argument to masrks variable
	marks := 80
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	switch {
	case grade == "A":
		fmt.Printf("Excellent!\n")
	case grade == "B", grade == "C":
		fmt.Printf("Well done\n")
	case grade == "D":
		fmt.Printf("You passed\n")
	case grade == "F":
		fmt.Printf("Better try again\n")
	default:
		fmt.Printf("Invalid grade\n")
	}
	fmt.Printf("Your grade is  %s\n", grade)
	fmt.Println("____________")

}

func EvenOdd() {
	for i := 1; i <= 4; i++ {
		if i%2 == 1 {
			fmt.Println(i, "odd")
		} else {
			fmt.Println(i, "even")
		}
	}
	fmt.Println("____________")

}
func FizzBuzz() {
	for i := 1; i <= 15; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		}
	}
	fmt.Println("____________")

}

func SelectOperator() {
	var c1, c2, c3 chan int
	var i1, i2 int

	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
	fmt.Println("____________")

}
