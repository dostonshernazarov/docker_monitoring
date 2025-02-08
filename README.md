Here’s a comprehensive `README.md` file for your project. It explains the purpose of the project, how to set it up, and how to use it. You can copy and paste this into a `README.md` file in your project's root directory.

---

# Docker Container Monitoring System

This project is a **Docker Container Monitoring System** that allows you to monitor the status of Docker containers by pinging their IP addresses and displaying the results in a web interface. The system consists of four main components:

1. **Backend Service**: A RESTful API written in Go that handles database operations.
2. **Frontend Service**: A React-based web application that displays the status of Docker containers.
3. **Pinger Service**: A Go service that pings Docker containers and sends the results to the Backend API.
4. **PostgreSQL Database**: Stores the status of Docker containers.
5. **Database migrations**: For avto migrate tables

The system is containerized using Docker and orchestrated with Docker Compose. Nginx is used as a reverse proxy to route requests between the frontend and backend.

---

## Table of Contents

1. [Features](#features)
2. [Technologies Used](#technologies-used)
3. [Prerequisites](#prerequisites)
4. [Setup and Installation](#setup-and-installation)
5. [Running the Project](#running-the-project)
6. [Project Structure](#project-structure)
7. [API Endpoints](#api-endpoints)
9. [Troubleshooting](#troubleshooting)
10. [Contributing](#contributing)
11. [License](#license)

---

## Features

- **Real-time Monitoring**: The Pinger service continuously pings Docker containers and updates the database.
- **Web Interface**: The Frontend service provides a user-friendly interface to view the status of Docker containers.
- **RESTful API**: The Backend service exposes endpoints to fetch and insert container status data.
- **Containerized**: All services are containerized using Docker for easy deployment and scalability.

---

## Technologies Used

- **Backend**: Go (Golang)
- **Frontend**: React (JavaScript)
- **Database**: PostgreSQL
- **Pinger Service**: Go (Golang)
- **Containerization**: Docker, Docker Compose
- **Reverse Proxy**: Nginx
- **Migrations**: SQL migrations

---

## Prerequisites

Before running the project, ensure you have the following installed:

1. **Docker**: [Install Docker](https://docs.docker.com/get-docker/)
2. **Docker Compose**: [Install Docker Compose](https://docs.docker.com/compose/install/)
3. **Node.js** (for local frontend development): [Install Node.js](https://nodejs.org/)
4. **Go** (for local backend development): [Install Go](https://golang.org/doc/install)

---

## Setup and Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/dostonshernazarov/docker_monitoring.git
   cd docker_monitoring
   ```

2. **Set Up Environment Variables**:
   - Create a `.env` file in the root directory with the following content:
     ```env
     POSTGRES_USER=postgres
     POSTGRES_PASSWORD=yourpassword
     POSTGRES_DB=docker_monitor
     ```

3. **Build and Run the Project**:
   ```bash
   docker-compose up --build
   ```

   This will:
   - Build Docker images for all services.
   - Start the PostgreSQL database with migrations, Backend, Frontend, Pinger, and Nginx services.

4. **Access the Application**:
   - Open your browser and navigate to `http://localhost:80`.
   - You should see the frontend interface displaying the status of Docker containers.

---

## Running the Project

### Using Docker Compose

To start all services, run:
```bash
docker-compose up --build
```

To stop all services, run:
```bash
docker-compose down
```

### Running Services Locally (Optional)

1. **Backend**:
   - Navigate to the `backend` directory:
     ```bash
     cd backend
     ```
   - Run the backend service:
     ```bash
     go run main.go
     ```

2. **Frontend**:
   - Navigate to the `frontend` directory:
     ```bash
     cd frontend
     ```
   - Install dependencies and start the frontend:
     ```bash
     npm install
     npm start
     ```

3. **Pinger**:
   - Navigate to the `pinger` directory:
     ```bash
     cd pinger
     ```
   - Run the pinger service:
     ```bash
     go run main.go
     ```

---

## Project Structure

```
docker-monitor/
├── backend/               # Backend service (Go)
│   ├── main.go            # Backend entry point
│   ├── Dockerfile         # Dockerfile for backend
│   └── go.mod             # Go dependencies
├── frontend/              # Frontend service (React)
│   ├── src/               # React source code
│   ├── Dockerfile         # Dockerfile for frontend
│   └── package.json       # Node.js dependencies
├── migrations/
|   ├── 00...table.down.sql
|   ├── 00...table.up.sql
├── pinger/                # Pinger service (Go)
│   ├── main.go            # Pinger entry point
│   ├── Dockerfile         # Dockerfile for pinger
│   └── go.mod             # Go dependencies
├── nginx.conf             # Nginx configuration
├── docker-compose.yml     # Docker Compose configuration
├── .env                   # Environment variables
└── README.md              # Project documentation
```

---

## API Endpoints

### Backend API

- **GET `/status`**:
  - Fetches the status of all Docker containers.
  - Example Response:
    ```json
    [
      {
        "ip_address": "192.168.1.1",
        "ping_time": 100,
        "last_success": "2023-10-01T12:00:00Z"
      }
    ]
    ```

- **POST `/status`**:
  - Inserts a new container status into the database.
  - Example Request Body:
    ```json
    {
      "ip_address": "192.168.1.1",
      "ping_time": 100,
      "last_success": "2023-10-01T12:00:00Z"
    }
    ```

---

## Troubleshooting

1. **CORS Issues**:
   - Ensure the Backend service has CORS enabled (see `backend/main.go`).
   - Use the proxy configuration in the Frontend (`frontend/package.json`).

2. **Database Connection Issues**:
   - Verify that the PostgreSQL service is running.
   - Check the `.env` file for correct database credentials.

3. **Pinger Service Not Sending Data**:
   - Check the logs of the Pinger service for errors.
   - Ensure the Backend API is running and accessible.

---

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Submit a pull request with a detailed description of your changes.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---
