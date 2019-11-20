package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// type NameReader interface {
// 	Read(source string, key string) string
// }

type ConsoleData struct {
	Filename string
	Key      string
}

func main() {
	var v ConsoleData
	v.loadData()

	fileContent := getDataFromFile(v.Filename)

	jsonData := parseFile(fileContent)
	fmt.Println(jsonData[v.Key])
}

func parseFile(data string) map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal([]byte(data), &result)
	return result
}

func (v *ConsoleData) loadData() {
	v.Filename = os.Args[1]
	v.Key = os.Args[2]
}

func getDataFromFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	checkError(err)
	return string(data)
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
