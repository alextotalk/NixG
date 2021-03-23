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
func swap2(x, y *int) {
	*x, *y = *y, *x
}
func EvenOdd(num int) (int, bool) {
	n := int(num / 2)
	return n, num%2 == 0
}

func maxNum(list ...int) (maxNum int) {
	for _, vol := range list {
		if maxNum < vol {
			maxNum = vol
		}
	}
	return
}
func makeOddGen() func() uint {
	i := uint(1)

	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fib2(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		if i < 2 {
			result[i] = i + 1
		} else {
			result[i] = result[i-1] + result[i-2]
		}
	}
	return result
}

func recoveryFunction() {
	if recoveryMessage := recover(); recoveryMessage != nil {
		fmt.Println(recoveryMessage)
	}
	fmt.Println("This is recovery function...")
}

func executePanic() {
	defer recoveryFunction()
	panic("This is Panic Situation")
	fmt.Println("The function executes Completely")
}

func main() {
	fmt.Println(fib(11))
	fmt.Println(fib2(11))

	nextMakeOddGen := makeOddGen()
	fmt.Println(nextMakeOddGen())
	fmt.Println(nextMakeOddGen())

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

	els := []int{
		48, 96, 86,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17, 111,
	}
	fmt.Println(maxNum(els...))

	executePanic()
	fmt.Println("Main block is executed completely...")

	defer second()
	first()

	defer func() {
		str1 := recover()
		fmt.Println(str1)
	}()

	panic("PANIC")
}
