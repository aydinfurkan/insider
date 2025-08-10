# Insider - Message Service

A Go-based message service with PostgreSQL and Redis support, containerized with Docker.

## Prerequisites

- [Docker](https://www.docker.com/get-started) (version 20.10 or higher)
- [Docker Compose](https://docs.docker.com/compose/install/) (version 2.0 or higher)

## Getting Started with Docker Compose

### 1. Clone the Repository

```bash
git clone <repository-url>
cd insider
```

### 2. Create Environment File

Create a `.env` file and fill it using `.env.example`

### 3. Build and Run the Application

```bash
# Build and start all services
docker-compose up --build

# Or run in detached mode (background)
docker-compose up --build -d
```

### Stopping Services

```bash
# Stop all services
docker-compose down

# Stop services and remove volumes
docker-compose down -v

# Stop services and remove images
docker-compose down --rmi all
```

### 4. Swagger adress for local

```bash
http://localhost:3000/swagger/index.html
```
