package main

import (
	"database/sql"
	"log"
	"os/exec"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error
	connStr := "host=monitor_postgres port=5432 user=postgres dbname=docker_monitor sslmode=disable password=yourpassword"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		pingContainers()
		time.Sleep(10 * time.Second)
	}
}

func pingContainers() {
	// Replace with actual Docker container IPs
	ips := []string{"192.168.1.1", "192.168.1.2"}

	for _, ip := range ips {
		cmd := exec.Command("ping", "-c", "1", ip)
		err := cmd.Run()
		pingTime := 0
		if err == nil {
			pingTime = 100 // Example ping time
		}

		_, err = db.Exec("INSERT INTO container_status (ip_address, ping_time, last_success) VALUES ($1, $2, $3)",
			ip, pingTime, time.Now())
		if err != nil {
			log.Println("Error inserting status:", err)
		}
	}
}
