package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const UserAgent = "Golang Reddit Reader"

func main() {
	// Create a request and add the proper headers.
	req, err := http.NewRequest("GET", "https://api.reddit.com/r/aww", nil)
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

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	for i := 1; i < len(response.Data.Children); i++ {
		currentItem := response.Data.Children[i].Data
		isImgurImg := strings.Contains(currentItem.URL, "imgur.com")
		isJpg := strings.Contains(currentItem.URL, ".jpg")
		if isImgurImg && isJpg {
			fmt.Println(currentItem.URL)
			downloadFile("C:/Dev/_images/img.jpg", currentItem.URL)
			idx := strings.Index(currentItem.URL, "imgur.com/")
			fmt.Println(idx)

		}

	}

	defer resp.Body.Close()

}

type Response struct {
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

func downloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
