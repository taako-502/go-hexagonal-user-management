package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Item struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	resp, err := http.Get("http://qiita.com/api/v2/users/snaka/items?page=1&per_page=10")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data []Item // nil slice
	// data := make([]Item, 0) のように要素数0の slice としても良い

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	for _, item := range data {
		fmt.Printf("%s %s\n", item.CreatedAt, item.Title)
	}
}
