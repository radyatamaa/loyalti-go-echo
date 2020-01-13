package router

import (
	"github.com/labstack/echo"
	"github.com/radyatamaa/loyalti-go-echo/src/api"
	"github.com/radyatamaa/loyalti-go-echo/src/api/apiKafka/Merchant"
	"github.com/radyatamaa/loyalti-go-echo/src/api/host"
	"github.com/radyatamaa/loyalti-go-echo/src/api/middlewares"
	"net/http"
)

func New() *echo.Echo {
	// create groups
	e := echo.New()
	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")
	jwtGroup := e.Group("/jwt")

	// set all middlewares
	middlewares.SetMainMiddlewares(e)
	middlewares.SetAdminMiddlewares(adminGroup)
	middlewares.SetCookieMiddlewares(cookieGroup)
	middlewares.SetJwtMiddlewares(jwtGroup)
	middlewares.SetCorsMiddlewares(e)

	// set main routes
	api.MainGroup(e)

	// set group routes
	api.AdminGroup(adminGroup)
	api.CookieGroup(cookieGroup)
	api.JwtGroup(jwtGroup)

	e.GET("/ping", Ping)
	e.POST("/merchant", Merchant.PublishMerchant)

	host.StartKafka()
	return e
}

func Ping(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
