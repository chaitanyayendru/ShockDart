# ShockDart

ShockDart is a powerful, scalable notification hub designed to manage multiple channels for sending notifications, including Firebase, AWS SES (for email), and SMS. This project is built using Go, Kafka, and PostgreSQL, with a flexible architecture that allows the easy addition of new notification channels.

---

## Table of Contents

- [Prerequisites](#prerequisites)
- [Project Structure](#project-structure)
- [Environment Variables](#environment-variables)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Docker Setup](#docker-setup)
- [Contributing](#contributing)

---

## Prerequisites

- **Golang** (>=1.18)
- **Docker** (for containerization)
- **Kafka** and **PostgreSQL** (configured through Docker Compose)

---

## Project Structure

```plaintext
ShockDart/
│
├── cmd/
│   └── main.go               # Main entry file to initialize and run the app
│
├── config/                
│   └── config.go              # Configuration file loading environment variables
│
├── internal/              
│   ├── api/               
│   │   ├── handler/       
│   │   │   └── notification_handler.go  # Handles API routes for notifications
│   │   └── router.go         # Router configuration
│   ├── service/              
│   │   └── notification_service.go      # Business logic for notifications
│   ├── repository/           
│   │   ├── kafka_repo.go     # Kafka integration for queuing notifications
│   │   └── pg_repo.go        # PostgreSQL repository for persistence
│   └── model/                
│       └── notification.go   # Model for notifications
│
├── pkg/                       
│   ├── logger.go              # Logging setup
│   ├── oauth.go               # OAuth2 configuration and helper functions
│   └── response.go            # Custom response formatting
│
├── Dockerfile                 # Docker configuration for the application
├── docker-compose.yml         # Docker Compose setup for the entire stack
├── go.mod                     # Go module dependencies
└── README.md                  # Project documentation
```

---

## Environment Variables

Create a `.env` file in the project root to define the necessary environment variables. Here’s an example:

```plaintext
SERVER_PORT=:8080
KAFKA_BROKER=localhost:9092
POSTGRES_URL=postgres://postgres:password@localhost:5432/shockdart?sslmode=disable
FIREBASE_PROJECT_ID=<firebase-project-id>

OAUTH_CLIENT_ID=<oauth-client-id>
OAUTH_CLIENT_SECRET=<oauth-client-secret>
OAUTH_REDIRECT_URL=http://localhost:8080/auth/callback
```

---

## Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/chaitanyayendru/shockdart.git
   cd shockdart
   ```

2. **Install Go Dependencies**:
   ```bash
   go mod tidy
   ```

3. **Build the Application**:
   ```bash
   go build -o shockdart cmd/main.go
   ```

---

## Usage

1. **Run the application** directly:
   ```bash
   ./shockdart
   ```

2. **Or, use Docker** to start the application along with Kafka and PostgreSQL:
   ```bash
   docker-compose up --build
   ```

---

## API Endpoints

### Authentication
- **[GET] /auth/login** - Initiates OAuth login process.
- **[GET] /auth/callback** - Handles OAuth callback to obtain the access token.

### Notifications
- **[POST] /api/notification** - Send a notification to the specified channel (Firebase, AWS SES, SMS).

#### Example Notification Request
```json
POST /api/notification
Content-Type: application/json

{
  "channel": "firebase",
  "message": "Hello, this is a test notification",
  "target": "device_token"
}
```

---

## Docker Setup

Docker is configured to handle dependencies like Kafka, PostgreSQL, and the application. Here’s a guide to the Docker commands and services provided in `docker-compose.yml`:

### Build and Run
```bash
docker-compose up --build
```

### Shut Down
```bash
docker-compose down
```

### Common Docker Commands

- **View Logs**:
  ```bash
  docker-compose logs -f
  ```
- **Access Database**:
  ```bash
  docker exec -it shockdart_postgres psql -U postgres
  ```

---

## Contributing

1. Fork the project.
2. Create a feature branch (`git checkout -b feature/YourFeature`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Open a pull request.

---

This README provides a comprehensive overview, setup instructions, and example API usage to help developers start working with ShockDart. Let me know if you need further customizations!