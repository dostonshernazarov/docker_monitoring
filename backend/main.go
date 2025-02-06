package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error
	connStr := fmt.Sprintf("postgres://%s:%s@db/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/containers", getContainers).Methods("GET")

	http.Handle("/", r)
	log.Println("Backend started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getContainers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT ip, last_ping FROM containers")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var ip string
		var lastPing string
		rows.Scan(&ip, &lastPing)
		fmt.Fprintf(w, "IP: %s, Last Ping: %s\n", ip, lastPing)
	}
}
