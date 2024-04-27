package apiserver

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"myapp/internal/app/models"
	"myapp/internal/storage"

	"github.com/labstack/echo/v4"
)

type Server struct {
	postgres *storage.PostgreSQL
}

func NewServer(postgres *storage.PostgreSQL) *Server {
	return &Server{postgres: postgres}
}

func (s *Server) GetBook(c echo.Context) error {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
	}

	query := "SELECT name, author FROM book WHERE id = $1"
	row := s.postgres.DB.QueryRow(query, bookId)
	if row.Err() != nil {
		return c.String(http.StatusBadRequest, "this book doesn't exist")
	}

	var Name sql.NullString
	var Author sql.NullString
	err = row.Scan(&Name, &Author)
	if err != nil {
		return c.String(http.StatusBadRequest, "this book doesn't exist")
	}
	book := models.NewBook(bookId, Name.String, Author.String)
	return c.JSON(http.StatusOK, book)
}

func (s *Server) GetAllBooks(c echo.Context) error {

	var m []models.Book
	query := "SELECT COUNT(*) FROM book"
	row := s.postgres.DB.QueryRow(query)
	if row.Err() != nil {
		return c.String(http.StatusBadRequest, "there is no books")
	}

	var count int
	err := row.Scan(&count)
	if err != nil {
		return c.String(http.StatusBadRequest, "error")
	}

	for i := range count {
		query = "SELECT name, author FROM book WHERE id = $1"
		row = s.postgres.DB.QueryRow(query, i+1)

		var Name sql.NullString
		var Author sql.NullString
		err = row.Scan(&Name, &Author)
		if err != nil {
			continue
		}
		book := models.NewBook(i+1, Name.String, Author.String)
		m = append(m, *book)

	}
	return c.JSON(http.StatusOK, m)
}

func (s *Server) PostBook(c echo.Context) error {
	book := models.Book{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&book)
	if err != nil {
		log.Printf("Failed processing postBook request: %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	query := "INSERT INTO book ( name, author ) VALUES ($1, $2) ON CONFLICT DO NOTHING"
	result, err := s.postgres.DB.Exec(query, book.Name, book.Author)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return c.String(http.StatusBadRequest, "can`t create book (or book already exists)")
	}
	return c.String(http.StatusOK, "we got your book")
}

func (s *Server) UpdateBook(c echo.Context) error {
	bookId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.String(http.StatusBadRequest, "enter valid id")
	}

	book := models.Book{}

	defer c.Request().Body.Close()

	err = json.NewDecoder(c.Request().Body).Decode(&book)
	if err != nil {
		log.Printf("Failed processing updateBook request: %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	query := "UPDATE book SET name = $1, author = $2 where id = $3"
	_, err = s.postgres.DB.Exec(query, book.Name, book.Author, bookId)
	if err != nil {
		return c.String(http.StatusBadRequest, "this book doesn't exist")
	}

	return c.String(http.StatusOK, "we update your book!")
}

func (s *Server) DeleteBook(c echo.Context) error {
	bookId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.String(http.StatusBadRequest, "enter valid id")
	}

	query := "DELETE  FROM book where id = $1"
	_, err = s.postgres.DB.Exec(query, bookId)

	if err != nil {
		return c.String(http.StatusBadRequest, "this book doesn't exist")
	}

	return c.String(http.StatusOK, "we delete your book!")
}
