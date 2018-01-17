package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"

	"github.com/tidwall/gjson"
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

func howToParse() {
	data := readInputFile("input.txt")
	blockData := gjson.Get(string(data), "block.data")
	fmt.Println("got blockData", blockData)
	fmt.Printf("type of blockdata is %T\n", blockData)
	fmt.Println("is array?", blockData.IsArray())
	blockData.ForEach(func(key, value gjson.Result) bool {
		fmt.Println("value is array?", value.IsArray())
		fmt.Println("data hash", value.Array()[0], "now hash value", value.Array()[1])
		return true
	})
}

func main() {
	//checkSha256()
	//orderCheck("{a:1,b:1}")
	//orderCheck("{b:1,a:1}")
	howToParse()
}
