package mw

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var AuthMiddleware = middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {

	cookie := &http.Cookie{}

	cookie.Name = "sessionID"
	cookie.Value = username + password
	cookie.Expires = time.Now().Add(48 * time.Hour)

	c.SetCookie(cookie)
	return true, c.String(http.StatusOK, "you were authorized!")
})

func CheckCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")

		if err != nil {
			log.Println(err)
			return err
		}

		if cookie.Value == "admin" {
			log.Println("right cookie")
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "you are not admin")
	}
}
