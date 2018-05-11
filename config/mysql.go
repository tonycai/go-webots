package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tkanos/gonfig"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"
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
	// insert
	start := time.Now()
	values1 := []string{"u_" + start.Format("2006-01-02 15:04:05") + ""}
	sql1 := "INSERT INTO w_user(user_name) VALUES(?);"
	last_id := insert_db(db, sql1, values1)
	fmt.Println("last_id: ", last_id)

	fmt.Println(reflect.TypeOf(db))

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

	// insert
	//stmt, err1 := db.Prepare("INSERT w_user SET user_name=?")
	/*
		stmt, err1 := db.Prepare("INSERT INTO w_user(user_name) VALUES(?);")
		if err1 == nil {
			res, err2 := stmt.Exec("alexa01")
			if err2 == nil {
				id, err3 := res.LastInsertId()
				if err3 == nil {
				}
				fmt.Println("last_id: ", id)
			}

		}
	*/

}
func insert_db(db *sql.DB, sql string, values []string) int64 {
	var last_id int64
	stmt, err1 := db.Prepare(sql)
	if err1 == nil {
		res, err2 := stmt.Exec(values[0])
		if err2 == nil {
			id, err3 := res.LastInsertId()
			if err3 == nil {
				last_id = id
			}
		}
	}
	return last_id
}
