package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func average1(xs ...float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	fmt.Println(x)

	return x * factorial(x-1)
}

func f() (int, int) {
	return 5, 6
}

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}
func swap(x *int, y *int) {
	var temp int
	temp = *x // save the value at address x
	*x = *y   // put y into x
	*y = temp // put temp into y
}
func EvenOdd(num int) (int, bool) {
	n := int(num / 2)
	return n, num%2 == 0
}
func main() {

	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
	fmt.Println(average1(xs...))
	fmt.Println(factorial(5))
	fmt.Println(f())
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1))
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	var a int = 100
	var b int = 200
	fmt.Printf("Before swap, value of a : %d\n", a)
	fmt.Printf("Before swap, value of b : %d\n", b)
	/* calling a function to swap the values.
	 * &a indicates pointer to a ie. address of variable a and
	 * &b indicates pointer to b ie. address of variable b.
	 */
	swap(&a, &b)
	fmt.Printf("After swap, value of a : %d\n", a)
	fmt.Printf("After swap, value of b : %d\n", b)
	fmt.Println(EvenOdd(2))
	fmt.Println(EvenOdd(1))
	fmt.Println(EvenOdd(1))

	defer second()
	first()

	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	panic("PANIC")
}
