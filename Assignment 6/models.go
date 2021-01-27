package mi=ai
import {
	"database/sql"
	_"github.com/go-sql-driver/mysql"
    "github.com/goriila/mux"
    "net/http"
}
type Instruction struct {
	ID string 'json:"id"'
	Name string 'json:"name"'
}
var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "<root>": <"JeyaMalaRavi@>@tcp(127.0.1.1:3306 / <db>")
	if(err!=nil)
	{
		panic(err.Error())
	} 
	defer db.close()

	router  := mux.NewRouter()

	router.HandleFunc("/posts", getPost).Methods("GET")
	router.HandleFunc("/posts", createPost). Methods("POST")
	router.HandleFunc("/posts {id}",getPost).Methods("GET")
	router.HandleFunc("/posts {id}",updatePost).Methods("PUT")
	router.HandleFunc("/posts {id}", deletePost).Methods("DELETE")
	
	http.ListenAndServe (":8000", router)
}