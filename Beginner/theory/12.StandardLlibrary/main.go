package main

import (
	"bytes"
	"container/list"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	fmt.Println(
		// true
		strings.Contains("test", "es"),
		"\n",

		// 2
		strings.Count("test", "t"),
		"\n",
		// true
		strings.HasPrefix("test", "te"),
		"\n",
		// true
		strings.HasSuffix("test", "st"),
		"\n",
		// 1
		strings.Index("test", "e"),
		"\n",
		// "a-b"
		strings.Join([]string{"a", "b"}, "-"),
		"\n",
		// == "aaaaa"
		strings.Repeat("a", 5),
		"\n",
		// "bbaa"
		strings.Replace("aaaa", "a", "b", 2),
		"\n",
		// []string{"a","b","c","d","e"}
		strings.Split("a-b-c-d-e", "-"),
		"\n",
		// "test"
		strings.ToLower("TEST"),
		"\n",
		// "TEST"
		strings.ToUpper("test"),
	)

	//Чтобы преобразовать строку в набор байт (и наоборот),
	//выполните следующие действия:

	arr := []byte("test")
	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println("it`s byte-", arr, "\n", "it`s string-", str)

	var buf bytes.Buffer
	fmt.Println(buf.Write([]byte("test")))
	s := buf.Bytes()
	fmt.Println(string(s))
	strings.NewReader("buf")
	fmt.Println(strings.NewReader("buf"))

	fmt.Println(readFile())
	fmt.Println(readFile())
	shotReadFile()
	shotCreatFile()
	openDir()
	openDirRecurs()
	fmt.Println(errorNew())
	containerList()
	sortData()
	hashesAndCryptography()
	checkHash()
	cryptHash()
	//testServer()
	//testServer2()
	RPC()
}

func readFile() error {
	file, err := os.Create("Beginner/theory/12.StandardLlibrary/test2.txt")
	if err != nil {
		// здесь перехватывается ошибка
		return err
	}

	file.WriteString("test")
	file, err = os.Open("Beginner/theory/12.StandardLlibrary/test.txt")
	if err != nil {
		// здесь перехватывается ошибка
		return err
	}
	defer file.Close()

	// получить размер файла
	stat, err := file.Stat()
	if err != nil {
		return err
	}
	// чтение файла
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return err
	}

	str := string(bs)
	fmt.Println(str)
	fmt.Println("ok")
	return nil
}
func shotReadFile() error {
	bs, err := ioutil.ReadFile("Beginner/theory/12.StandardLlibrary/test.txt")
	if err != nil {
		return err
	}
	str := string(bs)
	fmt.Println(str)
	return err
}
func shotCreatFile() error {
	file, err := os.Create("Beginner/theory/12.StandardLlibrary/shotCreatFile.txt")
	if err != nil {
		// здесь перехватывается ошибка
		return err
	}
	defer file.Close()

	file.WriteString("test")
	fmt.Println("File created")
	return err

}
func openDir() error {
	dir, err := os.Open(".")
	if err != nil {
		return err
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return err
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
	return err

}

func openDirRecurs() error {
	filepath.Walk("./Beginner/theory/12.StandardLlibrary/",
		func(path string, info os.FileInfo, err error) error {
			fmt.Println(path)
			return nil
		})
	return errors.New("error")
}

func errorNew() error {
	err := errors.New("error message!!!")
	return err

}

func containerList() {
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}

type Person struct {
	Name string
	Age  int
}

type ByName []Person

func (this ByName) Len() int {
	return len(this)
}
func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}

func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func sortData() {
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)
}
