package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type keyValue struct {
	Key   string
	Value string
}

type block struct {
	Nonce string `json:"nonce"`
	Data  []keyValue `json:"data"`
}
type iJson struct {
	Difficulty int   `json:"difficulty"`
	Block      block `json:"block"`
}

func inputJson(s []byte) string {
	// need to return json without white spaces
	// and keys need to be in alpha order!
	var ij iJson
	json.Unmarshal(s, &ij)
	fmt.Println("ij", ij)
	return ""
}

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
