package Controller
import (
	"echo_with_mysql_mvc/models"
	"net/http"
	"github.com/labstack/echo"
)
func GetEmployees (c echo.Context) error {
	result := models.GetEmployee()
	println("Expected Result")
	return c.JSON(http.StatusOK,result)
}