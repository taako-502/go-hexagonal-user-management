package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

func GetPSIPerformance() {

	r := gin.Default()
	r.GET("/get-psi-performance", func(context *gin.Context) {

		url := context.Query("url")

		body, err := run_api(url)
		if err != nil {
			log.Fatal(err)
		}

		var data Item
		if err := json.Unmarshal(body, &data); err != nil {
			log.Fatal(err)
		}

		context.JSON(200, gin.H{
			"id":    data.Id,
			"score": data.LighthouseResult.Categories.Performance.Score,
		})
	})
	log.Fatal(r.Run())
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
