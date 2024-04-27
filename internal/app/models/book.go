package models

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func NewBook(id int, name, author string) *Book {
	return &Book{
		Id:     id,
		Name:   name,
		Author: author,
	}
}
