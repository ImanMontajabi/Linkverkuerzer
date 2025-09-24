# Linkverkuerzer 🔗

A simple and fast URL shortener service built with Go and Fiber.

## Features

- ✅ Shorten long URLs
- ✅ Custom short codes
- ✅ Click tracking
- ✅ URL statistics
- ✅ Duplicate URL detection
- ✅ SQLite database storage

## Quick Start

### Prerequisites
- Go 1.25+ installed
- Git

### Installation

```bash
git clone https://github.com/ImanMontajabi/Linkverkuerzer.git
cd Linkverkuerzer
go mod tidy
go run .
```

The server will start at `http://localhost:3000`

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Health check |
| POST | `/shorten` | Create short URL |
| GET | `/:code` | Redirect to original URL |
| GET | `/stats/:code` | Get URL statistics |
| GET | `/urls` | List all URLs |

## Usage Examples

### Shorten a URL
```bash
curl -X POST http://localhost:3000/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://www.google.com"}'
```

### Access shortened URL
```bash
curl -L http://localhost:3000/abc123
```

### Get statistics
```bash
curl http://localhost:3000/stats/abc123
```

## Configuration

Set environment variables:

```bash
export PORT=3000
export DATABASE_URL=./urls.db
export BASE_URL=http://localhost:3000
export SHORT_CODE_LENGTH=6
export MAX_URL_LENGTH=2048
```

## Project Structure

```
├── main.go      # Application entry point
├── handlers.go  # HTTP request handlers
├── service.go   # Business logic
├── models.go    # Data structures
├── database.go  # Database operations
├── config.go    # Configuration management
└── utils.go     # Helper functions
```

## License

MIT License
