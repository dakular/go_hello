package main

import "fmt"
import "github.com/tidwall/gjson"
import "github.com/json-iterator/go"
import "github.com/buger/jsonparser"

const json = `{"name":{"first":"Tiger","last":"Liu"},"age":39}`

func main() {
	// var a int = 10
	// fmt.Printf("变量的地址: %x\n", &a)

	last_name := gjson.Get(json, "name.last").String()
	first_name := gjson.Get(json, "name.first").String()
	age := gjson.Get(json, "age").Int()
	fmt.Printf("%s %s is %d years old\n", first_name, last_name, age)

	last_name, err := jsonparser.GetString([]byte(json), "name", "last")
	first_name, err = jsonparser.GetString([]byte(json), "name", "first")
	age, err = jsonparser.GetInt([]byte(json), "age")
	if err == nil {
		fmt.Printf("%s %s is %d years old\n", first_name, last_name, age)
	}

	last_name = jsoniter.Get([]byte(json), "name", "last").ToString()
	first_name = jsoniter.Get([]byte(json), "name", "first").ToString()
	age = jsoniter.Get([]byte(json), "age").ToInt64()
	fmt.Printf("%s %s is %d years old\n", first_name, last_name, age)
}
