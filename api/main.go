package main

import (
	"fmt"
	"net/http"

	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		router := spinhttp.NewRouter()
		router.DELETE("/api/items/:id", deleteOne)
		router.POST("/api/items", addNew)
		router.GET("/api/items", getAll)
		router.ServeHTTP(w, r)
	})
}

func addNew(w http.ResponseWriter, r *http.Request, _ spinhttp.Params) {
	fmt.Fprintf(w, "This is the add new endpoint")
}

func getAll(w http.ResponseWriter, r *http.Request, _ spinhttp.Params) {
	fmt.Fprintf(w, "This is the get all endpoint")
}

func deleteOne(w http.ResponseWriter, r *http.Request, params spinhttp.Params) {
	fmt.Fprintf(w, "This is the deleteOne endpoint, you want to delete:  %s", params.ByName("id"))
}
func main() {}
