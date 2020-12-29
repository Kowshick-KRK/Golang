package main
import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)
type Article Struct
{
Title string 'json:"Title"'
Desc string  'json:"Desc"'
Content string 'json: "Content"'
}
type Articles [] Article

func allArticles (w http.Responsewrites,r * http.Request)
{
	articles := articles
	{
		Article{Title:"TestTitle",Desc:"Test Description",Content: "Hello World"}
	}
	fmt.Println("Endpoint Hit : All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}
func home Page(w http.ResponseWrites, r * http.Request)
{
	fmt.Fprintf(w,"News Updates in articles")
}

func handle Requests()
{
	http.handleFunc("/",homepage)
	http.handleFunc("/articles",all articles)
	log.Fatal(http.ListenAndServe("8081",nil))
}
func main()
{
	handle Request()
}
