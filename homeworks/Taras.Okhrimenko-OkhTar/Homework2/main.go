package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// NameReader -
// source can be local file or url of JSON file
// read this file, parse and return "key" field value
type NameReader interface {
	Read(source string, key string) string
}

// Use JSONPlaceholder (https://jsonplaceholder.typicode.com/)
// https://jsonplaceholder.typicode.com/users

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
	Street  string
	Suite   string
	City    string
	Zipcode string
	Geo     map[string]string
}

type Company struct {
	Name        string
	CatchPhrase string
	Bs          string
}

func Read(source string, key string) string {
	resp, err := http.Get(source)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var usersList []User
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, &usersList); err != nil {
		panic(err)
	}

	for _, elem := range usersList {
		fmt.Println("ID = ", elem.key)
	}

	return "1"

}

func main() {
	url := "https://jsonplaceholder.typicode.com/users"
	res := Read(url, "Name")

	fmt.Println(res)

}
