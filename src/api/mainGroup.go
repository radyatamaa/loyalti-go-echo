package api

import (
	"github.com/labstack/echo"
	"github.com/radyatamaa/loyalti-go-echo/src/api/handlers"
)

func MainGroup(e *echo.Echo) {
	e.GET("/login", handlers.Login)
	e.GET("/yallo", handlers.Yallo)
	e.POST("/graphql", handlers.QuerySong)
}
