package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"time"
)

type ContainerStatus struct {
	IPAddress   string    `json:"ip_address"`
	PingTime    int       `json:"ping_time"`
	LastSuccess time.Time `json:"last_success"`
}

func main() {
	for {
		pingContainers()
		time.Sleep(10 * time.Second) // Ping every 10 seconds
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

		status := ContainerStatus{
			IPAddress:   ip,
			PingTime:    pingTime,
			LastSuccess: time.Now(),
		}

		sendStatusToBackend(status)
	}
}

func sendStatusToBackend(status ContainerStatus) {
	jsonData, err := json.Marshal(status)
	if err != nil {
		log.Println("Error marshaling status:", err)
		return
	}

	resp, err := http.Post("http://backend:8080/status", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error sending status to backend:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Println("Backend returned non-201 status code:", resp.StatusCode)
	}
}
