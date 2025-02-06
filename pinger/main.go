package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	backendURL := "http://backend:8080/containers"
	for {
		resp, err := http.Get(backendURL)
		if err != nil {
			log.Println("Failed to fetch containers:", err)
		} else {
			log.Println("Fetched container data successfully")
		}
		resp.Body.Close()
		time.Sleep(30 * time.Second)
	}
}
