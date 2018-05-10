package main

// w_user
/*
 select * from w_user;
 created by tony
*/
import (
	"fmt"
	"github.com/tkanos/gonfig"
	"os"
)

type Configuration struct {
	Port     int
	Host     string
	Dbname   string
	User     string
	Password string
	Charset  string
}

func main() {

	configuration := Configuration{}
	err := gonfig.GetConf("./database.json", &configuration)
	if err != nil {
		os.Exit(500)
	}
	fmt.Println(configuration.Port)
	fmt.Println(configuration.Host)
	fmt.Println(configuration)

}
