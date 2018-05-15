package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/tkanos/gonfig"
	"io"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"regexp"
	"sync"
)

var (
	maxRoutineNum = 10
)

type Hospital struct {
	Name string
	Link string
}

type Configuration struct {
	Port     int
	Host     string
	Dbname   string
	User     string
	Password string
	Charset  string
	Leveldb  string
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func parser_html(file1 string, db *leveldb.DB) {

	b, err := ioutil.ReadFile(file1) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	base_url := "https://www.guahao.com"
	content := string(b) // convert content to a 'string'

	/*
	   <h1 style="display: inline-block; font-size: 20px; margin-right: 5px;">^M
	                   <strong><a style="display: inline-block;color:#000;line-height: 24px; font-weight: 100" href="/hospital/17634310-b567-4d21-8a02-20dc15e90da5000">北京协和医院</a></strong>^M
	               ^M
	               <h3 style="display: inline-block;">^M
	                   <span class="h3">^M
	                   三级甲等^M
	                   </span>^M
	               </h3>^M
	               </h1>^M
	*/
	//reg := regexp.MustCompile(`(?i:^hello).*Go`)
	//reg := regexp.MustCompile(`(?mi)<h1[\s\S]*?</h1>`) //忽略大小，非贪婪匹配
	reg := regexp.MustCompile(`(?mis)<h1.+?href="(/hospital.+?)">(.+?)</a>.+?</h1>`) //忽略大小，非贪婪匹配
	//fmt.Printf("%q\n", reg.FindAllString(str, -1))
	elements := reg.FindAllStringSubmatch(content, -1)

	for i, element := range elements {
		fmt.Println(i, base_url+element[1], element[2])
		fmt.Println("===========")
		hospital := &Hospital{Name: element[2], Link: element[1]}
		hos, err := json.Marshal(hospital)
		if err != nil {
			fmt.Println(err)
			return
		}
		uuid := GetMD5Hash(base_url + element[1])
		fmt.Println(string(uuid))
		fmt.Println(string(hos))
		save_json(db, string(uuid), string(hos))
	}
	/*
		match := reg.FindAllStringSubmatch(content, -1) //FindAllStringSubmatch会将捕获到的放到子slice
		if match != nil {
			fmt.Print("%#v", match)
		}
	*/

}
func get_dbfname() string {
	configuration := Configuration{}
	err := gonfig.GetConf("./config/database.json", &configuration)
	if err != nil {
		os.Exit(500)
	}
	return configuration.Leveldb
}

func save_json(db *leveldb.DB, k string, v string) bool {
	// put data
	err := db.Put([]byte(k), []byte(v), nil)
	//
	if err != nil {
		return false
	}
	return true
}
func main() {
	var wg sync.WaitGroup
	ch := make(chan int, maxRoutineNum) //maxRoutineNum = 10

	dbfname := get_dbfname()
	db, err := leveldb.OpenFile(dbfname, nil)
	if err != nil {
		log.Fatal("openFile: " + dbfname + " error!")
	}
	fmt.Println(reflect.TypeOf(db))
	defer db.Close()
	reader := bufio.NewReader(os.Stdin)

	for {
		file1, _, err := reader.ReadLine() // ReadRune
		if err != nil && err == io.EOF {
			break
		}
		wg.Add(1)
		ch <- 1
		if len(ch) >= maxRoutineNum+1 {
			fmt.Println("## ch满了, 处于阻塞")
		}
		go func(file string) {
			defer wg.Done()
			fmt.Println("Parser: ", file)
			parser_html(file, db)
			<-ch
		}(string(file1))
	}
	wg.Wait()
	fmt.Println("Done")
}
