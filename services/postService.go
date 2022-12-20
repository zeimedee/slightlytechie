package services

import (
	"database/sql"

	"github.com/zeimedee/slightlytechie/database"
)

func AddPost(title, body string) (*sql.Rows, error) {
	db := database.ConnectDb()
	sqlStatement := "INSERT INTO posts (title, body, created_at, updated_at) VALUES($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
	rows, err := db.Query(sqlStatement, title, body)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return rows, nil
}

func ReadAllPost() (*sql.Rows, error) {

	db := database.ConnectDb()
	sqlStatement := "SELECT * FROM posts"
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return rows, nil
}

func ReadPost(id string) (*sql.Rows, error) {
	db := database.ConnectDb()
	sqlStatement := "SELECT * FROM posts WHERE id = $1 LIMIT 1"

	rows, err := db.Query(sqlStatement, id)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	return rows, nil
}

func UpdatePost(id, title, body string) (map[string]string, error) {
	db := database.ConnectDb()
	sqlStatement := "UPDATE posts SET title = $1, body=$2, updated_at = CURRENT_TIMESTAMP WHERE id = $3"
	_, err := db.Query(sqlStatement, title, body, id)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return map[string]string{"message": "post updated"}, nil
}

func DeletePost(id string) (map[string]string, error) {
	db := database.ConnectDb()
	sqlStatement := "DELETE FROM posts WHERE id = $1"

	_, err := db.Query(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return map[string]string{"message": "post deleted"}, nil
}
