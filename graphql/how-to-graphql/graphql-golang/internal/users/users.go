package users

import (
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	database "github.com/tatsuro-m/hackernews/internal/pkg/db/migrations/mysql"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u User) Create() {
	statement, err := database.Db.Prepare("INSERT INTO Users(Username,Password) VALUES(?,?)")
	fmt.Println(statement)
	if err != nil {
		log.Fatalln(err)
	}

	hashedPassword, err := HashPassword(u.Password)
	_, err = statement.Exec(u.Username, hashedPassword)
	if err != nil {
		log.Fatalln(err)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetAll() []User {
	stmt, err := database.Db.Prepare("SELECT * FROM Users")
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	var resultUsers []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			log.Fatalln(err)
		}
		resultUsers = append(resultUsers, user)
	}

	return resultUsers
}

func GetUserIdByUsername(username string) (int, error) {
	statement, err := database.Db.Prepare("select ID from Users WHERE Username = ?")
	if err != nil {
		log.Fatalln(err)
	}
	row := statement.QueryRow(username)

	var id int
	err = row.Scan(&id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err)
		}
		return 0, err
	}

	return id, nil
}

func (u *User) Authenticate() bool {
	statement, err := database.Db.Prepare("select Password from Users WHERE Username = ?")
	if err != nil {
		log.Fatalln(err)
	}
	row := statement.QueryRow(u.Username)

	var hashedPassword string
	err = row.Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatalln(err)
		}
	}

	return CheckPasswordHash(u.Password, hashedPassword)
}
