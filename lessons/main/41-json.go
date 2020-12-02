package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type response1 struct {
	Page   int
	Fruits []string
}

// you can use struct tags to define exact field names
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// marshalling = go -> json conversion

	// primitive boolean conversion
	bolB, _ := json.Marshal(true)
	// by default json.Marshall return []byte which can be converted to string
	fmt.Println(bolB, string(bolB)) // [116 114 117 101] true

	// primitive int conversion
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB)) // 1

	// primitive float conversion
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB)) // 2.34

	// primitive string conversion
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB)) // "gopher"

	// array conversion
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB)) // ["apple","peach","pear"]

	// map conversion
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB)) // {"apple":5,"lettuce":7}

	// by default structs are exported with field names
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B)) // {"Page":1,"Fruits":["apple","peach","pear"]}

	// we can overwrite field names with `json` tags
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B)) // {"page":1,"fruits":["apple","peach","pear"]}

	// unmarshalling = json -> go conversion

	// unmarshalling to map string -> interface{}
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)                                    // map[num:6.13 strs:[a b]]
	fmt.Println(reflect.TypeOf(dat["num"]), dat["num"]) // float64 6.13

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1) // a

	// unmarshalling to struct string -> interface{}
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)           // {1 [apple peach]}
	fmt.Println(res.Fruits[0]) // apple

	// example of writing encoded json directly to io.Writer without using intermediate string
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
