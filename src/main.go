package main

import (
	"database/sql"
	"example.com/m/v2/src/db"
	"example.com/m/v2/src/interfunc"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

var DB *sql.DB

func main() {
	db.InitDB()
	defer DB.Close()
	r := mux.NewRouter()
	r.HandleFunc("/query", interfunc.Query).Methods(http.MethodGet)
	r.HandleFunc("/add", interfunc.Add).Methods(http.MethodPost)
	r.HandleFunc("/delete", interfunc.Delete).Methods(http.MethodPost)
	r.HandleFunc("/change", interfunc.Change).Methods(http.MethodPost)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}
