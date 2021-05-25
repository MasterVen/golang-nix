package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	createWriteFile(string(bytes), id)
}

func createWriteFile(data string, id string) {

	c, err := os.Create("./filesData/" + id + ".txt")
	if err != nil {
		log.Fatal("whoops", err)
	}

	ws, err := c.WriteString(data)
	if err != nil {
		log.Fatal("whoops", err)
	}
	fmt.Println("bytes written", ws)

}
