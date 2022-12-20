package services

import (
	"database/sql"

	"github.com/zeimedee/slightlytechie/database"
)

func AddPost() error {
	return nil
}

func ReadAllPost() (*sql.Rows, error) {

	db := database.ConnectDb()
	sqlStatement := "SELECT * FROM posts"
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	db.Close()
	return rows, nil
}

func ReadPost(id string) (*sql.Rows, error) {
	db := database.ConnectDb()
	sqlStatement := "SELECT * FROM posts WHERE id = $1 LIMIT 1"

	rows, err := db.Query(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	db.Close()
	return rows, nil
}

func UpdatePost() error {
	return nil
}

func DeletePost() error {
	return nil
}
