package models

type Post struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	Created_at string `json:"created_at"`
}
