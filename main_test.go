package main

import "testing"

func Test_1_run_api(t *testing.T) {
	var url string = "https://media-engine.jp"
	_, err := run_api(url)
	if err != nil {
		t.Error(`miss`)
	}
}

func Test_2_run_api(t *testing.T) {
	var url string = "hogehogehoge" //存在しないURL
	_, err := run_api(url)
	if err != nil {
		t.Error(`miss`)
	}
}
