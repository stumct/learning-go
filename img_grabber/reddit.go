package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type RedditResponse struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

type Item struct {
	Title string
	URL   string
	Score int64
}

func (i *Item) hasImgurURL() bool {
	return strings.Contains(i.URL, "imgur.com")
}
func (i *Item) hasExtension(ext string) bool {
	return strings.Contains(i.URL, ext)

}

func GetSubreddit(url string) *Response {
	// Create a request and add the proper headers.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)
	// Handle the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		//return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		//return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("error reading body")
	}
	//fmt.Println(string(body))

	var response RedditResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	return &response
}
