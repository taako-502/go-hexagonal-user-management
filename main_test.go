package main

import "testing"

func testA(t *testing.T) {
	var url string = "https://media-engine.jp"
	_, err := run_api(url)
	if err != nil {
		t.Error(`miss`)
	}
}
