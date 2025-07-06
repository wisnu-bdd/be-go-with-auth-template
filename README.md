# Backend Go With Auth Template

A Go-based backend service template including JWT based authentication setup. This service provides authentication and (other future features) functionality through a RESTful API.

## Features

- JWT-based authentication
- Protected endpoints
- CORS support
- Environment-based configuration

## Prerequisites

- Go 1.24.3 or higher
- Environment variables setup (see Configuration section)

## Installation

1. Clone the repository:
```bash
git clone github.com/wisnu-bdd/be-go-with-auth-template
cd be-go-with-auth-template
```

2. Install dependencies:
```bash
go mod download
```

## Configuration

Create a `.env` file in the root directory with the following variables:

```env
PORT=8000
JWT_SECRET=yourJwtSecret
ALLOWED_ORIGINS=http://localhost:3000,http://your-frontend
```

## Available Endpoints

- `POST /login` - Authentication endpoint
  - Requires email and password in request body
  - Returns JWT token on successful authentication

- `GET /protected` - Protected test endpoint
  - Requires valid JWT token in Authorization header

## Development

To run the server locally:

```bash
go run main.go
```

The server will start on http://localhost:8080 (or your configured PORT)

## Building for Production

### Build for EC2 (Linux)

```bash
GOOS=linux GOARCH=amd64 go build -o server
```

### Build for Local Development

```bash
go build -buildvcs=false -o myapp
```

## Project Structure

```
.
├── config/       # Configuration and environment setup
├── handlers/     # HTTP request handlers
├── middleware/   # Middleware functions (CORS, JWT)
├── models/       # Data models
├── utils/        # Utility functions
├── main.go      # Application entry point
└── README.md    # This file
```

## Dependencies

- github.com/golang-jwt/jwt/v5 - JWT token handling
- github.com/joho/godotenv - Environment configuration

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b dev/new-feature`)
3. Commit your changes (`git commit -m 'Add some new feature'`)
4. Push to the branch (`git push origin dev/new-feature`)
5. Open a Pull Request

## License

<!-- This project is proprietary software of BDD SKS Digital. -->