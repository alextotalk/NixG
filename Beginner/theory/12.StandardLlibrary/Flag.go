package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func Flags() {
	// Определение флагов
	maxp := flag.Int("max", 6, "the max value")
	// Парсинг
	flag.Parse()
	// Генерация числа от 0 до max
	fmt.Println("Flag - some number", rand.Intn(*maxp))
}
