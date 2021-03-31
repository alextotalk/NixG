package main

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
	"io/ioutil"
)

func hashesAndCryptography() {
	h := crc32.NewIEEE()
	h.Write([]byte("some text test"))
	v := h.Sum32()
	fmt.Println(v, "-hashSum")
}
func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	h := crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(), nil
}
func checkHash() error {
	h1, err := getHash("Beginner/theory/12.StandardLlibrary/test.txt")
	if err != nil {
		fmt.Println(err)
		return err
	}
	h2, err := getHash("Beginner/theory/12.StandardLlibrary/test2.txt")
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(h1, h2, h1 == h2)
	return nil
}
func cryptHash() {
	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}
