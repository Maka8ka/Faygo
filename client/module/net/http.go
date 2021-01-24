package net

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return string("error")
	} else {
		defer resp.Body.Close()
		
		
		
		status := resp.Status
		
		return string(status)
	}

}

func HttpPost(url string, body string) string {
	resp, err := http.Post(url, "text/html", strings.NewReader(body))
	if err != nil {
		fmt.Println(err)
		return string("error")
	} else {
		defer resp.Body.Close()
		response, _ := ioutil.ReadAll(resp.Body)
		return string(response)
	}
}
