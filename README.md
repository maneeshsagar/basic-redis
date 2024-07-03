# Go Redis Project

This is a simple Go application that connects to a Redis instance using Docker Compose. The project demonstrates how to set up a Go application with Redis, using a Dockerized environment.

## Project Structure

basic-redis/
├── Dockerfile
├── docker-compose.yml
├── main.go
├── db/
│ └── redis.go
├── config/
│ └── config.go
└── handlers/
  └── handler.go 
└── service/
  └── putService.go
  ...
└── persistence/
  └── persistence.go 
  ...
└── util/
  └── response.go 
  ...
└── iolayer/
  └── request.go 
  
## Prerequisites

- Docker and Docker Compose installed on your machine.

## Setup and Run

1. **Clone the repository:**

   ```sh
   git clone https://github.com/maneeshsagar/basic-redis.git
   cd repository
   docker-compose up --build
  ```sh

   

  
