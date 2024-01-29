package helper

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	file = "foo.db"
)

func Load() *sql.DB {
	database, err := sql.Open("sqlite3", file)
	if err != nil {
		Setup()
	}

	return database
}

func Setup() {
	database, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Panic(err)
	}

	defer database.Close()

	sqlStmt := `
	create table teams (id integer not null primary key, name text);
	create table players (id integer not null primary key, name text);
	create table match (id integer not null primary key, name text);
	create table series (id integer not null primary key, name text);
	create table matchplayer (id integer not null primary key, name text);
	`

	_, err = database.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	// transaction, err := database.Begin()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// statement, err := transaction.Prepare("insert into foo(id, name) values(?, ?)")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer statement.Close()
	// for i := 0; i < 100; i++ {
	// 	_, err = statement.Exec(i, fmt.Sprintf("こんにちは世界%03d", i))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// err = transaction.Commit()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
