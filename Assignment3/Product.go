package main
import(
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"

)
func get(w.httpResponceWriter,r * http.Request)
{
w.Header().Set("Product-Type",)
w.WriteHeader(http.StatusCreated)
w.Write([]byte(`{"message": "get product"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Product-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "post Product}`))
}

func put(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Product-Type", "application/json")
    w.WriteHeader(http.StatusAccepted)
    w.Write([]byte(`{"message": "put Product"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Product-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "delete Product"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Product-Type", "application/json")
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte(`{"message": "not found"}`))
}
func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", get).Methods(http.MethodGet)
    r.HandleFunc("/", post).Methods(http.MethodPost)
    r.HandleFunc("/", put).Methods(http.MethodPut)
    r.HandleFunc("/", delete).Methods(http.MethodDelete)
    r.HandleFunc("/", notFound)
	log.Fatal(http.ListenAndServe(":8080", r))
}