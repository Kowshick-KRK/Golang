package db
import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

//Creating postgres DB Cnnection
func CreateCon() *sql.DB{
	db,err := sql.Open("postgres", "root: @tcp(localhost:3306/test")
	if err!=nil
	{
		fmt.Println(err.Error())
	}
	else
	{
		fmt.Println("DB is Connected")
	}
	err := b.Ping()
	fmt.Println(err)
	if err!=nil{
		fmt.Println("DB is not Connected")
		fmt.Println(err.Error())
	}
	return db
}