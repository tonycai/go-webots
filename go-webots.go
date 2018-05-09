package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	//"strconv"
	"strings"
	"sync"
	//"strings"
)

var i int

func uuid() string {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s", out)
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	var wg sync.WaitGroup
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if info.Mode()&os.ModeNamedPipe == 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | gocowsay")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		url, _, err := reader.ReadLine() // ReadRune
		if err != nil && err == io.EOF {
			break
		}
		wg.Add(1)
		i++
		go func(url string) {
			defer wg.Done()
			//var fn string = uuid()
			tokens := strings.Split(url, "/")
			fileName := tokens[len(tokens)-1]
			fmt.Println("Downloading", url, "to", fileName)
			//output, err := os.Create("./hospital/" + strconv.Itoa(i) + ".html") // strconv.Itoa(i)
			output, err := os.Create("./hospital/" + GetMD5Hash(url) + ".html") // strconv.Itoa(i)
			if err != nil {
				log.Fatal("Error while creating", fileName, "-", err)
			}
			defer output.Close()

			//res, err := http.Get(url)
			client := &http.Client{}

			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				log.Fatalln(err)
			}

			req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.367")

			resp, err := client.Do(req)

			if err != nil {
				log.Fatal("http get error: ", err)
			} else {
				defer resp.Body.Close()
				_, err = io.Copy(output, resp.Body)
				if err != nil {
					log.Fatal("Error while downloading", url, "-", err)
				} else {
					fmt.Println("Downloaded", fileName)
				}
			}
		}(string(url))
		fmt.Printf("%s\n", url)
	}
	wg.Wait()
	fmt.Println("Done")
}
