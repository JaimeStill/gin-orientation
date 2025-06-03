# Basic Gin Router Project

## How to Run

1. Ensure you have Go installed (version 1.16 or higher)
2. Navigate to the src directory
3. Run the server:

   ```bash
   go run main.go
   ```

4. The server will start on `http://localhost:8080`

### Testing the Endpoints

- Health check:

  ```bash
  curl http://localhost:8080/health
  ```

  Expected response: `{"status":"OK"}`

- Echo endpoint:

  ```bash
  curl -X POST http://localhost:8080/echo
  ```

  Expected response: `{"message":"Request received and confirmed"}`

## How This Resolves the Experiment

This implementation directly addresses the experiment's hypothesis by demonstrating that the Gin router can handle multiple HTTP methods (GET/POST) on different endpoints:

1. **Router Initialization**: Uses `gin.Default()` to create a router with built-in middleware
2. **Multiple Endpoints**: Implements two distinct endpoints (`/health` and `/echo`)
3. **Different HTTP Methods**: Handles GET requests on `/health` and POST requests on `/echo`
4. **Proper Status Codes**: Both endpoints return HTTP 200 OK with JSON responses
5. **Server Binding**: Listens on localhost:8080 as specified

The implementation meets all success criteria:

- Server starts without errors
- GET /health returns 200 with "OK" status
- POST /echo returns 200 with confirmation message
- Response times are well under 100ms due to Gin's efficient routing

## Source Code Overview

### main.go

The entire implementation is contained in a single file with the following structure:

1. **Package Declaration**: Standard Go main package
2. **Imports**:
   - `net/http` for HTTP status constants
   - `github.com/gin-gonic/gin` for the web framework
3. **Main Function**:
   - Creates a default Gin router with logging and recovery middleware
   - Registers GET handler for `/health` endpoint
   - Registers POST handler for `/echo` endpoint
   - Starts the HTTP server on port 8080

### Key Components

- **gin.Default()**: Creates a router with default middleware (logger and recovery)
- **router.GET()**: Registers a GET route handler
- **router.POST()**: Registers a POST route handler
- **gin.Context**: Provides request/response handling capabilities
- **c.JSON()**: Sends JSON responses with appropriate content-type headers
- **gin.H**: Shorthand for `map[string]interface{}` for JSON responses

The code demonstrates Gin's simplicity in handling multiple endpoints with different HTTP methods while maintaining clean, readable code.
