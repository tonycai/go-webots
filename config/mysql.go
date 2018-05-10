package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tkanos/gonfig"
	"log"
	"os"
	"strconv"
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
	var conn string
	conn = configuration.User + ":" + configuration.Password +
		"@tcp(" + configuration.Host + ":" + strconv.Itoa(configuration.Port) + ")" +
		"/" + configuration.Dbname + "?charset=" + configuration.Charset
	fmt.Println(conn)
	/*
		var iconn = []byte{'t', 3, 'b'}
		fmt.Println(iconn)
	*/
	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	//fmt.Println(configuration)

	var (
		id   int
		name string
	)

	rows, err := db.Query("select user_id, user_name from w_user where user_id >= ? order by user_id desc;", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() // Important to close! This is no-op if already closed

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

	err = rows.Err() // check at end of the loop
	if err != nil {
		log.Fatal(err)
	}
}
