package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	urls := []string{
		"http://example.com/2/file1",
		"http://example.com/2/file2",
		"http://example.com/2/file3",
	}

	var wg sync.WaitGroup

	wg.Add(len(urls))

	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			tokens := strings.Split(url, "/")
			fileName := tokens[len(tokens)-1]
			fmt.Println("Downloading", url, "to", fileName)

			output, err := os.Create("./hospital/" + fileName)
			if err != nil {
				log.Fatal("Error while creating", fileName, "-", err)
			}
			defer output.Close()

			res, err := http.Get(url)
			if err != nil {
				log.Fatal("http get error: ", err)
			} else {
				defer res.Body.Close()
				_, err = io.Copy(output, res.Body)
				if err != nil {
					log.Fatal("Error while downloading", url, "-", err)
				} else {
					fmt.Println("Downloaded", fileName)
				}
			}
		}(url)
	}
	wg.Wait()
	fmt.Println("Done")
}
