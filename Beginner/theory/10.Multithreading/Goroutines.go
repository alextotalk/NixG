package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}
func f2(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		atm := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * atm)
	}
}

func Goroutines1() {
	go f(0)
	var input string
	fmt.Scanln(&input)
}
func Goroutines10() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}
func GoroutineWithDuration() {
	for i := 0; i < 3; i++ {
		go f2(i)
	}
	var input string
	fmt.Scanln(&input)
}
