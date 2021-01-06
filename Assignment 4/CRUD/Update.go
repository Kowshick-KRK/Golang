package main
import (
	"net/http"
	"database/sql"
	"github.com/labstack/echo"
	"fmt"
	"log"
	_"github.com/lib/pq"
)
e.PUT("/employee", func(c echo.Context) error {
	u := new(Employee)
	if err := c.Bind(u); err != nil {
		return err
	}
	PlsqlStatement := "UPDATE employees SET name=$1,salary=$2,age=$3 WHERE id=$5"
	res, err := db.Query(PlsqlStatement, u.Name, u.Salary, u.Age, u.Id)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusCreated, u)
	}
	return c.String(http.StatusOK, u.Id)
})