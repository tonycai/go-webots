package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
)

var (
	maxRoutineNum = 10
)

type stringFlag struct {
	set   bool
	value string
}

func (sf *stringFlag) Set(x string) error {
	sf.value = x
	sf.set = true
	return nil
}

func (sf *stringFlag) String() string {
	return sf.value
}

var directory stringFlag

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

func init() {
	flag.Var(&directory, "dir", "The directory")
}

func main() {
	var local_files_dir string = "local_files"
	flag.Parse()
	if directory.set {
		local_files_dir = directory.value
	}
	var wg sync.WaitGroup
	ch := make(chan int, maxRoutineNum) //maxRoutineNum = 10
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
		ch <- 1
		if len(ch) >= maxRoutineNum+1 {
			fmt.Println("## ch满了, 处于阻塞")
		}
		go func(url string) {
			defer wg.Done()
			//var fn string = uuid()
			tokens := strings.Split(url, "/")
			fileName := tokens[len(tokens)-1]
			fmt.Println("Downloading", url, "to", fileName)
			output, err := os.Create("./" + local_files_dir + "/" + GetMD5Hash(url) + ".html") // strconv.Itoa(i)
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
			<-ch
		}(string(url))
		fmt.Printf("%s\n", url)
	}
	wg.Wait()
	fmt.Println("Done")
}
