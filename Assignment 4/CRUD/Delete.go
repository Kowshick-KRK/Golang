package main
import (
	"net/http"
	"database/sql"
	"github.com/labstack/echo"
	"fmt"
	"log"
	_"github.com/lib/pq"
)
id := c.Param("id")
	PlsqlStatement := "DELETE FROM employees WHERE id = $1"
	res, err := db.Query(PlsqlStatement, id)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusOK, "Deleted")
	}
	return c.String(http.StatusOK, id+"Deleted")
})