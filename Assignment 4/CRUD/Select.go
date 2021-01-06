package main
import (
	"net/http"
	"database/sql"
	"github.com/labstack/echo"
	"fmt"
	"log"
	_"github.com/lib/pq"
)
e.GET("/employee", func(c echo.Context) error {
	PlsqlStatement := "SELECT id, name, salary, age FROM employees order by id"
	rows, err := db.Query(PlsqlStatement)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	}
	defer rows.Close()
	result := Employees{}
 
	for rows.Next() {
		employee := Employees{}
		err2 := rows.Scan(&employee.Id, &employee.Name, &employee.Salary, &employee.Age)
		// Exit if we get an error
		if err2 != nil {
			return err2
		}
		result.Employees = append(result.Employees, Employee)
	}
	return c.JSON(http.StatusCreated, result)
 
	//return c.String(http.StatusOK, "ok")
})