# Go HTTP Server with QR Code Generation

This project is a simple HTTP server built with Go that provides an endpoint to generate QR codes from a given string. It uses the `chi` router for routing and the `github.com/skip2/go-qrcode` package for QR code generation.

## Project Structure

```
qrcode-backend
├── cmd
│   └── server
│       └── main.go        # Entry point of the application
├── internal
│   ├── handlers
│   │   └── qrcode.go      # Handler for generating QR codes
│   └── routes
│       └── routes.go      # Route definitions
├── go.mod                  # Module definition
├── go.sum                  # Dependency checksums
└── README.md               # Project documentation
```

## Endpoints

### POST /generate-qr

This endpoint accepts a JSON payload with a `data` field (string) and generates a QR code.

**Request Example:**

```json
{
  "data": "Hello, World!"
}
```

**Response Example:**

```json
{
  "qr_code": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAA..."
}
```

### GET /healthz

This endpoint returns a simple health check response.

**Response Example:**

```
200 OK
ok
```

## Getting Started

1. Clone the repository:

   ```
   git clone https://github.com/yourusername/go-http-server.git
   cd go-http-server
   ```

2. Install dependencies:

   ```
   go mod tidy
   ```

3. Run the server:

   ```
   go run cmd/server/main.go
   ```

4. Test the endpoints using a tool like `curl` or Postman.

## Dependencies

- `github.com/skip2/go-qrcode` for QR code generation
- `github.com/go-chi/chi` for routing

## License

This project is licensed under the MIT License. See the LICENSE file for details.
