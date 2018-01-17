package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"

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

type fakeKeyValue struct {
	key string
	val int
}

func testSort() {
	a := make([]fakeKeyValue, 0)
	a = append(a, fakeKeyValue{"213a7f85c41b256b679e530c4943eee0", -79})
	a = append(a, fakeKeyValue{"94a82aefc5f498f0d3538381ef728afa", 81})
	a = append(a, fakeKeyValue{"753c06c1e88d26b4d62f7edba70e9673", -85})
	fmt.Println(a)
	fmt.Println("now need to sort")
	sort.Slice(a, func(i, j int) bool {
		return a[i].key < a[j].key
	})
	fmt.Println("sorted")
	fmt.Println(a)
}

func combineParseAndSort() []fakeKeyValue {
	data := readInputFile("input.txt")
	blockData := gjson.Get(string(data), "block.data")
	a := make([]fakeKeyValue, 0)
	blockData.ForEach(func(key, value gjson.Result) bool {
		fmt.Println("value is array?", value.IsArray())
		k := value.Array()[0].String()
		v := value.Array()[1].Int()
		a = append(a, fakeKeyValue{k, int(v)})
		return true
	})
	//fmt.Println("gathered data:", a)
	sort.Slice(a, func(i, j int) bool {
		return a[i].key < a[j].key
	})
	//fmt.Println("sorted")
	//fmt.Println(a)
	return a
}

func makeString(a []fakeKeyValue, nonce int) string {
	ts := "{\"nonce\":" + strconv.Itoa(nonce) + ",\"data\":["
	for i := 0; i < len(a); i++ {
		ts += "[" + a[i].key + "," + strconv.Itoa(a[i].val) + "]"
		if i < len(a)-1 {
			ts += ","
		}
	}
	ts += "]"
	return ts
}

func makeDataLine() {
	a := combineParseAndSort()
	ts := makeString(a, 1000)
	fmt.Println(ts)
}

func main() {
	//checkSha256()
	//orderCheck("{a:1,b:1}")
	//orderCheck("{b:1,a:1}")
	//howToParse()
	//testSort()
	//combineParseAndSort()
	makeDataLine()
}
