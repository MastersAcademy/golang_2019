// NEW version

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// NameReader -
// source can be local file or url of JSON file
// read this file, parse and return "key" field value
type NameReader interface {
	Read(source string, key string) (string, error)
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}

// getJSONFromURLByte func for getting json from url
func getJSONFromURLByte(url string) ([]byte, error) {
	var myClient = &http.Client{Timeout: 10 * time.Second}

	r, err := myClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	respByte := buf.Bytes()

	return respByte, nil
}

func (u *User) Read(source string, key string) (string, error) {
	var v map[string]interface{}

	if err := json.Unmarshal([]byte(source), &v); err != nil {
		return "", err
	}

	if val, ok := v[key]; ok {
		return fmt.Sprintf("%v", val), nil
	}
	return "No key", nil
}

func main() {
	url := "https://jsonplaceholder.typicode.com/users/3"

	jsonStream, err := getJSONFromURLByte(url)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Json:\n %v\n", string(jsonStream))

	var usr User
	res, err := usr.Read(string(jsonStream), "name")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(res)

}
