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
	Source, Key string
}

func main() {
	var v ConsoleData
	v.loadData()

	fileContent := getData(v.Source)
	jsonData := parseData(fileContent)

	if jsonData[v.Key] == nil {
		fmt.Printf("Key '%s' dosn't exists.\n", v.Key)
	} else {
		fmt.Println(jsonData[v.Key])
	}
}

func parseData(data string) map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal([]byte(data), &result)
	return result
}

func (v *ConsoleData) loadData() {
	v.Source = os.Args[1]
	v.Key = os.Args[2]
}

func getData(source string) string {
	data, err := ioutil.ReadFile(source)
	checkError(err)
	return string(data)
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
