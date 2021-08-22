package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

type post struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
}

func main() {
	var arr = make([]int, 0)
	ch := make(chan int)

	for i := 1; i <= 200; i++ {
		go httpRequest(i, ch)
	}

	for i := 0; i < 200; i++ {
		arr = append(arr, <-ch)

	}
	close(ch)

	sort.Ints(arr)
	log.Println(arr)
}

// httpRequest membutuhkan waktu
// test 1 : 0m2,355s
// test 2 : 0m2,088s
// test 3 : 0m2,400s
func httpRequest(i int, ch chan<- int) {
	str := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", i)

	resp, err := http.Get(str)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalln(resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var data post
	json.Unmarshal(body, &data)

	ch <- data.ID
}
