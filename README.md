# IoT Device Server

A simple REST API server built with Go, Gin framework, and MySQL database for managing IoT devices.

## Features

- REST API endpoints for device management
- MySQL database integration using GORM
- Health check endpoint
- Environment variable configuration
- Auto-migration of database schemas

## Prerequisites

- Go 1.24.2 or later
- MySQL 8.0 (running via Docker Compose)

## Environment Variables

The server can be configured using the following environment variables:

- `DB_HOST`: Database host (default: localhost)
- `DB_PORT`: Database port (default: 3306)
- `DB_USER`: Database user (default: root)
- `DB_PASSWORD`: Database password (default: root)
- `DB_NAME`: Database name (default: iot-schema)
- `PORT`: Server port (default: 8080)

## Running the Server

1. Make sure your MySQL container is running:
   ```bash
   docker-compose up -d
   ```

2. Start the server:
   ```bash
   go run main.go
   ```

The server will start on port 8080 (or the port specified in the PORT environment variable).

## API Endpoints

### Health Check
- `GET /health` - Check if the server is running

### Device Management
- `GET /api/v1/devices` - Get all devices
- `POST /api/v1/devices` - Create a new device
- `GET /api/v1/devices/:id` - Get a device by ID

### Example Device JSON
```json
{
  "name": "Temperature Sensor",
  "type": "sensor"
}
```

## Testing the API

You can test the API using curl:

```bash
# Health check
curl http://localhost:8080/health

# Get all devices
curl http://localhost:8080/api/v1/devices

# Create a device
curl -X POST http://localhost:8080/api/v1/devices \
  -H "Content-Type: application/json" \
  -d '{"name": "Temperature Sensor", "type": "sensor"}'

# Get a specific device
curl http://localhost:8080/api/v1/devices/1
```