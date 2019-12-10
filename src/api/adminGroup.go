package api

import (
	"github.com/labstack/echo"
	"github.com/radyatamaa/loyalti-go-echo/src/api/handlers"
)

func AdminGroup(g *echo.Group) {
	g.GET("/main", handlers.MainAdmin)
}
