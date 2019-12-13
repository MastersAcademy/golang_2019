// OLD version

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"
)

// NameReader -
// source can be local file or url of JSON file
// read this file, parse and return "key" field value
type NameReader interface {
	Read(source string, key string) []string
}

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"username"`
	Email   string  `json:"email"`
	Addr    Addr    `json:"address"`
	Phone   string  `json:"phone"`
	WebSite string  `json:"website"`
	Company Company `json:"company"`
}

type Addr struct {
	Street  string            `json:"street"`
	Suite   string            `json:"suite"`
	City    string            `json:"city"`
	Zipcode string            `json:"zipcode"`
	Geo     map[string]string `json:"geo"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func Read(source string, key string) []string {
	r, err := myClient.Get(source)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var usersList []User
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, &usersList); err != nil {
		panic(err)
	}

	var valArr = make([]string, 0, len(usersList))

	for _, user := range usersList {
		e := reflect.ValueOf(&user).Elem()

		var resArr = make([]string, 0, e.NumField())

		for i := 0; i < e.NumField(); i++ {
			varName := e.Type().Field(i).Name
			resArr = append(resArr, varName)
		}

		for i, v := range resArr {
			if key == v {
				varValue := e.Field(i).Interface()
				valArr = append(valArr, fmt.Sprintf("%v", varValue))
			}
		}

	}

	return valArr
}

func main() {
	url := "https://jsonplaceholder.typicode.com/users"
	res := Read(url, "Name")

	fmt.Println(res)
}
