package net

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpGet(url string) string {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	response, _ := ioutil.ReadAll(resp.Body)
	return string(response)
}

func HttpPost(url string, body string) string {
	resp, _ := http.Post(url, "text/html", strings.NewReader(body))
	response, _ := ioutil.ReadAll(resp.Body)
	return string(response)
}
