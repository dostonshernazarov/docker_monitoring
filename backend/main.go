package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type ContainerStatus struct {
	IPAddress   string    `json:"ip_address"`
	PingTime    int       `json:"ping_time"`
	LastSuccess time.Time `json:"last_success"`
}

var db *sql.DB

func main() {
	var err error
	connStr := "host=monitor_postgres port=5432 user=postgres dbname=docker_monitor sslmode=disable password=yourpassword"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	// Enable CORS
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/status", getStatus).Methods("GET")
	router.HandleFunc("/status", addStatus).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT ip_address, ping_time, last_success FROM container_status")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var statuses []ContainerStatus
	for rows.Next() {
		var status ContainerStatus
		if err := rows.Scan(&status.IPAddress, &status.PingTime, &status.LastSuccess); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		statuses = append(statuses, status)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(statuses)
}

func addStatus(w http.ResponseWriter, r *http.Request) {
	var status ContainerStatus
	if err := json.NewDecoder(r.Body).Decode(&status); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO container_status (ip_address, ping_time, last_success) VALUES ($1, $2, $3)",
		status.IPAddress, status.PingTime, status.LastSuccess)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
