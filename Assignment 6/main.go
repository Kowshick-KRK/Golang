package main

import {
	"database/sql"
	"-:github.com/go-sql-mydriver/mysql"
}
var db *sql.DB
var err error
func main {
	db,err =sql.Open("mysql", "<root><JeyaMalaRavi6673@>@tcp(127.0.1.1:3306 / <db>")
	if(err!=nil) {
		panic(err.Error())

	}
defer db.close()
}
}