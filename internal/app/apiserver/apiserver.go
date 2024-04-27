package apiserver

import (
	"fmt"
	"myapp/internal/app/mw"

	"myapp/internal/storage"

	"github.com/labstack/echo/v4"
)

type App struct {
	s    *Server
	echo *echo.Echo
}

func New(postgres *storage.PostgreSQL) (*App, error) {
	a := &App{}

	a.s = NewServer(postgres)

	a.echo = echo.New()

	a.echo.GET("/books/:id", a.s.GetBook)
	a.echo.GET("/books", a.s.GetAllBooks)

	a.echo.POST("/books", a.s.PostBook, mw.CheckCookie)
	a.echo.PUT("/books/:id", a.s.UpdateBook, mw.CheckCookie)
	a.echo.DELETE("/books/:id", a.s.DeleteBook, mw.CheckCookie)

	g := a.echo.Group("/")
	g.Use(mw.AuthMiddleware)

	return a, nil
}

func (a *App) Start(port string) error {

	fmt.Println("server running")

	return a.echo.Start(port)
}
