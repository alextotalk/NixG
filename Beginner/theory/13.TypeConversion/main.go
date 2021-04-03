package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var index int8 = 15

	var bigIndex int32

	bigIndex = int32(index)

	fmt.Println(bigIndex)
	fmt.Printf("index data type:    %T\n", index)
	fmt.Printf("bigIndex data type: %T\n", bigIndex)

	var big int64 = 64

	var little int8

	little = int8(big)

	fmt.Println(little)

	var x int64 = 57

	var y float64 = float64(x)

	fmt.Printf("%.4f\n", y)

	var f float64 = 390.8
	var i int = int(f)

	fmt.Printf("f = %.2f\n", f)
	fmt.Printf("i = %d\n", i)

	a := 5 / 2
	fmt.Println("a =  ", a)
	a2 := 5.0 / 2
	fmt.Println("a2 =  ", a2)

	a3 := strconv.Itoa(12)
	fmt.Printf("%q\n", a3)

	user := "Sammy"
	lines := 50

	fmt.Println("Congratulations, " + user + "! You just wrote " + strconv.Itoa(lines) + " lines of code.")

	fmt.Println(fmt.Sprint(421.034))

	f2 := 5524.53
	fmt.Println("f2=" + fmt.Sprint(f2))

	lines_yesterday := "50"
	lines_today := "108"

	yesterday, err := strconv.Atoi(lines_yesterday)
	if err != nil {
		log.Fatal(err)
	}

	today, err := strconv.Atoi(lines_today)
	if err != nil {
		log.Fatal(err)
	}
	lines_more := today - yesterday

	fmt.Println("lines_more=", lines_more)

	ax := "my string"

	b := []byte(ax)

	c := string(b)

	fmt.Println(ax)

	fmt.Println(b)

	fmt.Println(c)

}
