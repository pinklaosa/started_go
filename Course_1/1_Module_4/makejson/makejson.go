package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var addr string

	fmt.Print("Enter name: ")
	fmt.Scan(&name)

	fmt.Print("Enter address: ")
	fmt.Scan(&addr)

	data := make(map[string]string)

	data["name"] = name
	data["addr"] = addr

	d,_ := json.Marshal(data)

	fmt.Println(string(d))

}
