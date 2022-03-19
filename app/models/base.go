package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"

	// "fmt"
	"log"
	"todo_app/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

//テーブルの作成のコードを書いていく

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	// defer Db.Close()

	cmdU := `CREATE TABLE IF NOT EXISTS ` + tableNameUser + ` ( 
		id INTEGER PRIMARY KEY AUTO_INCREMENT, 
		uuid varchar(100) NOT NULL UNIQUE, name varchar(100), 
		email varchar(100), 
		password varchar(100), 
		created_at DATETIME )`

	_, err = Db.Exec(cmdU)
	if err != nil {
		log.Fatalln(err)
	}

	// Db.Close()
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (crypttext string) {
	crypttext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return crypttext
}