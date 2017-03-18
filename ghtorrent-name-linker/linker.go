package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *sql.DB
var mongo *mgo.Database

const lookupsAtOnce = 50

type user struct {
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Login string `json:"login" bson:"login"`
}

func main() {
	var err error
	db, err = sql.Open("mysql", "user:pass@host/ghtorrent") // local mysql database with the ghtorrent sql dump
	if err != nil {
		panic(err)
	}
	mgoSession, err := mgo.Dial("mongodb://ghtorrentro:ghtorrentro@localhost/github") // Dial the ghtorrent mongodb over ssh
	if err != nil {
		panic(err)
	}
	mongo = mgoSession.DB("github")

	needsMore := true
	offset := 0
	max := getLargestID()

	for needsMore {
		doneChannel := make(chan bool)
		usernames := fetchUsers(lookupsAtOnce, offset)
		for _, username := range usernames {
			go addName(username, doneChannel)
		}

		if offset > max {
			needsMore = false
			fmt.Println("End")
		}
		fmt.Println("Importing", strconv.Itoa(len(usernames)))
		if len(usernames) > 0 {
			for i := 0; i < len(usernames); i++ {
				<-doneChannel
			}
		}

		offset += lookupsAtOnce

		fmt.Println("Next round", strconv.Itoa(offset))
	}
}

func fetchUsers(limit, offset int) []string {
	rows, err := db.Query("SELECT login FROM users WHERE id >= ? and id < ? AND name IS NULL", offset, offset+limit)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	c := 0
	names := []string{}
	for rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			log.Fatal(err)
		}
		names = append(names, username)
		c++
	}
	return names
}

func addName(login string, c chan bool) {
	name, _ := lookUpName(login)
	if name == "" {
		c <- true
		return
	}
	rows, _ := db.Query("UPDATE users SET name=? WHERE login=?", name, login)
	defer rows.Close()
	c <- true
}

func lookUpName(login string) (string, error) {
	session := mongo.Session.Copy()
	defer session.Close()

	c := mongo.C("users").With(session)
	result := user{}
	err := c.Find(bson.M{"login": login}).Sort("-updated_at").One(&result)
	if err != nil {
		return "", err
	}
	return result.Name, nil
}

func getLargestID() int {
	rows, err := db.Query("SELECT id FROM users ORDER BY id DESC LIMIT 1")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var num int
	rows.Next()
	rows.Scan(&num)

	return num
}
