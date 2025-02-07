# Receipt Processor

This is a receipt processing web service written in Go. It calculates points based on certain rules provided in the API specification.

## Running Locally

### Prerequisites
- Docker (to run in a container)
- Go (optional for running directly without Docker)

### Running with Docker

1. Clone the repository:
    ```bash
    git clone <https://github.com/bhargavkonakanchi/receipt-processor>
    cd receipt-processor
    ```

2. Build the Docker image:
    ```bash
    docker build -t receipt-processor .
    ```

3. Run the container:
    ```bash
    docker run -p 8080:8080 receipt-processor
    ```

### Running without Docker

1. Clone the repository:
    ```bash
    git clone <https://github.com/bhargavkonakanchi/receipt-processor>
    cd receipt-processor
    ```

2. Install Go if it's not installed already. Instructions: https://golang.org/doc/install

3. Run the Go application:
    ```bash
    go run main.go
    ```

4. Access the API at `http://localhost:8080`.
   Process Receipt API:
   ```bash
    curl -X POST http://localhost:8080/receipts/process \                          
    -H "Content-Type: application/json" \
    -d @receipt.json
   ```

  Get Points API:
   ```bash
   curl http://localhost:8080/receipts/{id}/points
   ```

### API Endpoints

#### 1. Process Receipt

**POST** `/receipts/process`
- **Description**: This endpoint processes the receipt returns the receiptID for the receipt.
  
Example Request Body:
```json
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },
    {
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    }
  ],
  "total": "35.35"
}
```

Example Response:
```json
{ "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }
```

#### 2. Get Points

**GET** `/receipts/{id}/points`
**Description**: This endpoint calculates the reward points and returns the number of points for the receiptID.

Example Response:
20


