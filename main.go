package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/taako-502/go-sample-api/api"
)

type Item struct {
	Id               string           `json:"id"`
	CaptchaResult    string           `json:"captchaResult"`
	LighthouseResult LighthouseResult `json:"lighthouseResult"`
}

type LighthouseResult struct {
	Categories Categories `json:"categories"`
}

type Categories struct {
	Performance Performance `json:"performance"`
}

type Performance struct {
	Score float64 `json:"score"`
}

func run_api(url string) ([]byte, error) {
	resp, err := http.Get("https://www.googleapis.com/pagespeedonline/v5/runPagespeed?url=" + url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func main() {
	body, err := run_api("https://hepere.com")
	if err != nil {
		log.Fatal(err)
	}

	var data Item
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v: %v\n", data.Id, data.LighthouseResult.Categories.Performance.Score)

	//api
	api.Hello()
}
