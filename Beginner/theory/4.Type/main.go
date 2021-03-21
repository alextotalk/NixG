package main

import "fmt"

func main() {
	var x float64 = 20.0

	y := 42
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)

	var a int = 4
	var b int32
	var c float32
	var ptr *int

	/* example of type operator */
	fmt.Printf("Line 1 - Type of variable a = %T\n", a)
	fmt.Printf("Line 2 - Type of variable b = %T\n", b)
	fmt.Printf("Line 3 - Type of variable c= %T\n", c)
	fmt.Printf("Line 4 - Type of variable ptr= %T\n", ptr)
	//* для разименования , aдрес памяти
	// & для назначения переменной

	/* example of & and * operators */
	ptr = &a /* 'ptr' now contains the address of 'a'*/
	fmt.Printf("value of a is  %d\n", a)
	fmt.Printf("ptr is (address) %d.\n", ptr)
	fmt.Printf("*ptr is %d.\n", *ptr)
	ans := (true && false) || (false && true) || !(false && false)
	//true
	sum := 32132 * 42452
	fmt.Println(sum, "\n", ans)
	//length number of symbol
	fmt.Println(len("Hello World"))
	//название в байтах
	fmt.Println("Hello World"[0])
	//конкатенация
	fmt.Println("Hello " + "World\n")
}
