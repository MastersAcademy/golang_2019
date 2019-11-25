package main

//type file struct {
//	name string
//}
//
//var results []map[string]string
//
//func (file) Read(sourse, key string) string {
//	usersJson, _ := ioutil.ReadFile(sourse)
//	json.Unmarshal([]byte(usersJson), &results)
//	for key, result := range results {
//		if result["Age"] == key {
//			return fmt.Sprint(key, result["name"], result["type"])
//		}
//	}
//	return "Not found"
//}
//
//func main() {
//	sourse := file{"homeworks/ivan.horbatko-psixona/homework2/json/users.json"}
//	key := "17"
//	file()
//
//}

//Users struct which contains
//an array of users
//type Users struct {
//	Users []User `json:"users"`
//}
//
////User struct which contains a name
////a type and a list of social links
//
//func main() {
//	// Open our jsonFile
//	userJson, err := os.Open("homeworks/ivan.horbatko-psixona/homework2/json/users.json")
//	// if we os.Open returns an error then handle it
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("Successfully Opened users.json")
//	// defer the closing of our jsonFile so that we can parse it later on
//	defer userJson.Close()
//	//data := Users{}
//	//_ = json.Unmarshal([]byte(userJson), &data)
//	//for i:=0; i < len(data.Users); i++ {
//	//	fmt.Println("name:",data.Users[i].Name)
//	//	fmt.Println("age:",data.Users[i].Age)
//	//}
//	usersJson, _ := ioutil.ReadFile("homeworks/ivan.horbatko-psixona/homework2/json/users.json")
//	json.Unmarshal([]byte(usersJson), &results)

//	}
//

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type NameReader interface {
	Read(source string, key string) string
}

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Age  int    `json:"age"`
}

var users Users

const adr = "homeworks/ivan.horbatko-psixona/homework2/users.json"

func Read(sourse, key string) []string {
	//Open JSON file
	userJson, err := os.Open(sourse)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Already opened file")
	//Don't forget to close file
	defer userJson.Close()

	byteValue, _ := ioutil.ReadAll(userJson)
	json.Unmarshal(byteValue, &users)
	var valList = make([]string, 0, len(users.Users))
	var valListName = make([]string, 0, len(users.Users))
	var valListType = make([]string, 0, len(users.Users))
	var valListAge = make([]string, 0, len(users.Users))

	for i := 0; i < len(users.Users); i++ {
		valListName = append(valListName, users.Users[i].Name)
		valListType = append(valListType, users.Users[i].Type)
		valListAge = append(valListAge, strconv.Itoa(users.Users[i].Age))
	}
	switch key {
	case "Name":
		valList = valListName
	case "Type":
		valList = valListType
	case "Age":
		valList = valListAge
	}
	return valList
}

func main() {
	res := Read(adr, "Name")
	fmt.Println(res)
	res = Read(adr, "Age")
	fmt.Println(res)

}
