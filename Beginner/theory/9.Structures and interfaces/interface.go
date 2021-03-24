package main

import (
	"fmt"
	"time"
)

type Animal interface {
	Speak() string
}
type Dog struct {
}

func (d Dog) Speak() string {
	return ("Wow!")
}

type Cat struct {
}

func (c Cat) Speak() string {
	return ("Meooowww!")
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "??????"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "Design Patterns"
}

func DoSomething(v interface{}) {
	// ...
}
func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

// start with a string representation of our JSON data
var input = `
{
	"created_at": "Thu May 31 00:00:01 +0000 2012"
}
`

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	v, err := time.Parse(time.RubyDate, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	*t = Timestamp(v)
	return nil
}

//
//func main() {
//	animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
//	for _, animal := range animals {
//		fmt.Println(animal.Speak())
//	}
//
//	names := []interface{}{"stanley", "david", "oscar"}
//	PrintAll(names)
//
//	// our target will be of type map[string]interface{}, which is a pretty generic type
//	// that will give us a hashtable whose keys are strings, and whose values are of
//	// type interface{}
//	var val map[string]Timestamp
//
//	if err := json.Unmarshal([]byte(input), &val); err != nil {
//		panic(err)
//	}
//
//	fmt.Println(val)
//	for k, v := range val {
//		fmt.Println(k, reflect.TypeOf(v))
//	}
//	fmt.Println(time.Time(val["created_at"]))
//}
