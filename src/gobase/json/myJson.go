package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"` //自定义 json格式数据的key 名称
	Fruits []string `json: "fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("nivelle")
	fmt.Println(string(strB))

	//slice
	strD := []string{"apple", "peach", "pear"}
	jsrD, _ := json.Marshal(strD)
	fmt.Println(string(jsrD))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	jsMapD, _ := json.Marshal(mapD)
	fmt.Println(string(jsMapD))

	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peacjh"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}

	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	// 这里的map[string] interface {} 将保存一个string 为键，值为任意值的map
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	str := `{"page":1,"fruits":["apple","peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), res)
	fmt.Println(res)
	fmt.Println(*res)
	fmt.Println(res.Fruits[1])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"aoole": 5, "lettuce": 7}
	enc.Encode(d)

}
