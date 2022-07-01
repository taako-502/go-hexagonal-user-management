package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Item struct {
	Id            string `json:"id"`
	CaptchaResult string `json:"captchaResult"`
}

func main() {
	resp, err := http.Get("https://www.googleapis.com/pagespeedonline/v5/runPagespeed?url=https://hepere.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data Item // nil slice
	//data := make([]Item, 0) のように要素数0の slice としても良い

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s: %s", data.Id, data.CaptchaResult)
}
