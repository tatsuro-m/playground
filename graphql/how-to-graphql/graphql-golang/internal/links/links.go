package links

import (
	"fmt"
	"log"

	database "github.com/tatsuro-m/hackernews/internal/pkg/db/migrations/mysql"
	"github.com/tatsuro-m/hackernews/internal/users"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func (link Link) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address,UserID) VALUES(?,?, ?)")
	if err != nil {
		log.Fatalln(err)
	}

	res, err := stmt.Exec(link.Title, link.Address, link.User.ID)
	if err != nil {
		log.Fatalln(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatalln("Error:", err.Error())
	}
	log.Println("Row inserted!")

	return id
}

func GetAll() []Link {
	stmt, err := database.Db.Prepare("select L.id, L.title, L.address, L.UserID, U.Username from Links L inner join Users U on L.UserID = U.ID")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var links []Link
	var username string
	var id string

	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address, &id, &username)
		if err != nil {
			log.Fatal(err)
		}
		link.User = &users.User{ID: id, Username: username}
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return links
}

func DeleteByID(id string) bool {
	statement, err := database.Db.Prepare("DELETE FROM Links WHERE ID=?")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(id)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
