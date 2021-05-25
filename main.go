package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	countPost(100)
}

func countPost(num int) {
	wg := new(sync.WaitGroup)
	for i := 0; i < num; i++ {
		go post(strconv.Itoa(i), wg)
	}
	wg.Wait()
}

func post(id string, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	requestUrl := "https://jsonplaceholder.typicode.com/posts/" + id

	resp, err := http.Get(requestUrl)
	if err != nil {
		panic(err)
	}
	bytes, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		panic(errRead)
	}
	fmt.Println(string(bytes))
}
