package processToken

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func RouterProcessToken (c echo.Context) error{
	fmt.Println("masuk ke router proces Token")
	resp := ProcessToken()
	fmt.Println("Respon", resp)

	return c.JSON(http.StatusOK, resp)
}
