package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	file, err := ioutil.ReadFile("./data.json")
	if err != nil {
		fmt.Print(err)
		return
	}

	var records []Record

	err = json.Unmarshal(file, &records)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	recordMap := make(map[string]Record)

	for _, record := range records {
		if record.isInvalid() {
			fmt.Println(record.Id)
		} else {
			var key = record.getKey()
			r, found := recordMap[key]

			if found {
				fmt.Println(record.Id)
				fmt.Println(r.Id)
			} else {
				recordMap[key] = record
			}
		}
	}
}
