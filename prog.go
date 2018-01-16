package main

import (
	"crypto/sha256"
	"fmt"
)

func inputJson(s string) string {
	// need to return json without white spaces
	// and keys need to be in alpha order!
	return s
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

func main() {
	//checkSha256()
	orderCheck("{a:1,b:1}")
	orderCheck("{b:1,a:1}")

}
