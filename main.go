package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	requestUrl := "https://jsonplaceholder.typicode.com/posts/"

	resp, err := http.Get(requestUrl)
	if err != nil {
		panic(err)
	}
	bytes, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		panic(errRead)
	}
	fmt.Print(string(bytes))

}
