package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const UserAgent = "Golang Reddit Reader"

func main() {
	/*response := GetSubreddit("https://api.reddit.com/r/pics?limit=100")
	posts := response.Data.Children

	var wg sync.WaitGroup
	wg.Add(len(posts))
	for i := 1; i < len(posts); i++ {
		go func(id int) {
			post := posts[id].Data
			if post.hasImgurURL() && post.hasExtension(".jpg") {
				filename := post.URL[strings.Index(post.URL, "imgur.com/")+len("imgur.com/"):]
				downloadFile("C:/Dev/_images/"+filename, post.URL)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()*/

	response := GetImgurAPI("https://api.imgur.com/3/gallery/r/aww/", "abd05f05b578afe")
	fmt.Println(response)

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
