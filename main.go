package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	log, err := read("log.txt")
	if err != nil {
		panic(err)
	}
	report := parse(log)
	b, errjson := json.Marshal(report)
	if errjson != nil {
		fmt.Println(errjson)
		return
	}
	fmt.Println(string(b))
}
