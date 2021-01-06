package main
import (
	"net/http"
	"database/sql"
	"github.com/labstack/echo"
	"fmt"
	"log"
	_"github.com/lib/pq"
)
e.POST("/employee", func(c echo.Context) error {
	u := new(Employee)
		if err := c.Bind(u); err != nil {
			return err
		}
		PlsqlStatement := "INSERT INTO employees (name, salry,age)VALUES ($1, $2, $3)"
		res, err := db.Query(PlsqlStatement, u.Name, u.Salary, u.Age)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, u)
		}
		return c.String(http.StatusOK, "ok")
	})