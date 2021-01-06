package models
import (
	_"database/sql"
	"echo_with_mysql_mvc/db"
	"fmt"
	"database/sql"
)
type Employee struct {
	Id     string `json:"id"`
	Name   string `json:"employee_name"`
	Salary string `json: "employee_salary"`
	Age    string `json : "employee_age"`
}
type Employees struct {
	Employees []Employee `json:"employee"`
}
var con *sql.DB
 
func GetEmployee() Employees {
	con := db.CreateCon()
	//db.CreateCon()
	PlsqlStatement := "SELECT id,employee_name, employee_age, employee_salary FROM employee order by id"
 
	rows, err := con.Query(PlsqlStatement)
	fmt.Println(rows)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	}
	defer rows.Close()
	result := Employees{}
 
	for rows.Next() {
		employee := Employee{}
 
		err2 := rows.Scan(&employee.Id, &employee.Name, &employee.Salary, &employee.Age)
		// Exit if we get an error
		if err2 != nil {
			fmt.Print(err2)
		}
		result.Employees = append(result.Employees, employee)
	}
	return result
 
}