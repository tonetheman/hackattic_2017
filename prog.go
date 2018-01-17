package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
)

func checkSha256() {
	s := sha256.New()
	s.Write([]byte("{\"data\":[],\"nonce\":45}"))
	res := s.Sum(nil)
	// res is a byte array
	fmt.Printf("%x\n", res)
}

func orderCheck(s string) {
	h := sha256.New()
	h.Write([]byte(s))
	res := h.Sum(nil)
	// res is a byte array
	fmt.Printf("%x\n", res)
}

func readInputFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("got an error reading file", err)
	}
	return data
}

func main() {
	//checkSha256()
	//orderCheck("{a:1,b:1}")
	//orderCheck("{b:1,a:1}")
	data := readInputFile("input.txt")
	inputJson(data)
}
