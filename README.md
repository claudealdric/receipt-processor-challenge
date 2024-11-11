# Receipt Processor

A Go service that processes receipts and calculates reward points based on
receipt data.

## Prerequisites

- [Go 1.23](https://go.dev/dl/) or higher
- Git

## Running the application

1. Clone the repository:

   ```bash
   git clone https://github.com/claudealdric/receipt-processor-challenge.git
   cd receipt-processor-challenge
   ```

   **Note**: No additional dependencies are required. This project only uses the
   Go standard library.

2. Start the server:

   ```bash
   go run main.go
   ```

   By default, the server will start on `http://localhost:8080`.

## Project Structure

- `api/`: HTTP handlers, routing, and server logic
- `assert/`: Test utility functions
- `data/`: Data storage implementations
- `types/`: Type definitions and models
- `main.go`: Application entry point

## API Endpoints

### Process Receipt

- Endpoint: `/receipts/process`
- HTTP method: POST

```bash
curl -X POST http://localhost:8080/receipts/process \
  -H "Content-Type: application/json" \
  -d '{
    "retailer": "Target",
    "purchaseDate": "2022-01-01",
    "purchaseTime": "13:01",
    "items": [
      {
        "shortDescription": "Mountain Dew 12PK",
        "price": "6.49"
      }
    ],
    "total": "6.49"
  }'
```

### Get Points

- Endpoint: `/receipts/{id}/points`
- HTTP method: GET

```bash
curl http://localhost:8080/receipts/{id}/points
```
