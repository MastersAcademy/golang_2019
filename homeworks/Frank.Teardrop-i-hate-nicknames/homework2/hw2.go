package main

import(
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "encoding/json"
)

const NOT_FOUND = "key not found"

type NameReader interface {
    Read(source string, key string) string
}

type JSONFetcher map[string]string

func (fetcher JSONFetcher) Read(source string, key string) string {
    if val, ok := fetcher.retreive(source, key); ok {
        return val
    }
    data, err := readFromSource(source)
    if err != nil {
        panic(err)
    }
    var decoded map[string]interface{}
    err = json.Unmarshal(data, &decoded)
    if err != nil {
        panic(err)
    }
    value := findByKey(decoded, key)
    if (key != NOT_FOUND) {
        fetcher.store(source, key, value)
    }
    return value
}

func (fetcher JSONFetcher) store(source string, key string, value string) {
    fetcher[source+key] = value
}

func (fetcher JSONFetcher) retreive(source string, key string) (string, bool) {
    val, ok := fetcher[source+key]
    return val, ok
}

func readFromSource(source string) ([]byte, error) {
    if _, err := url.ParseRequestURI(source); err == nil {
        res, err := http.Get(source)
        if err != nil {
            return nil, err
        }
        defer res.Body.Close()
        return ioutil.ReadAll(res.Body)
    } else {
        return ioutil.ReadFile(source)
    }
}

func findByKey(decoded map[string]interface{}, key string) string {
    val, ok := decoded[key]
    if !ok {
        return NOT_FOUND
    }
    return val.(string)
}

func main() {
    fetcher := make(JSONFetcher)
    fmt.Println(fetcher.Read("https://jsonplaceholder.typicode.com/posts/1", "body"))
    fmt.Println(fetcher.Read("test.json", "key1"))
}