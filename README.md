# Receipt Processor

## Overview

This project implements a receipt processing service that awards points based on specific rules. It provides two endpoints:

1. **Process Receipt**: Submits a receipt for processing and returns an ID.
2. **Get Points**: Retrieves the points awarded for a given receipt ID.

## Requirements

- Go 1.18+
- Docker (optional, for containerized deployment)

## Build and Run

### Using Go

1. **Build the Application**

    ```sh
    go build -o main.exe ./cmd
    ```

2. **Run the Application**

    ```sh
    ./main.exe
    ```

### Using Docker

1. **Build the Docker Image**

    ```sh
    docker build -t reciept-processor-fetch .
    ```

2. **Run the Docker Container**

    ```sh
    docker run -p 8080:8080 reciept-processor-fetch
    ```

## API Endpoints

### Process Receipt

- **URL:** `/receipts/process`
- **Method:** `POST`
- **Payload:** JSON object representing the receipt
- **Response:** JSON object with the ID of the receipt

#### Example Request

```sh
curl -X POST http://localhost:8080/receipts/process -d '{
  "retailer": "Target",
  "purchaseDate": "2022-01-02",
  "purchaseTime": "13:13",
  "total": "1.25",
  "items": [
    {"shortDescription": "Pepsi - 12-oz", "price": "1.25"}
  ]
}' -H "Content-Type: application/json"
```
#### Example Response

```json
{
  "id": "some-unique-id"
}
```
### Get Points

- **URL:** `/receipts/{id}/points`
- **Method:** `GET`
- **Response:** JSON object with the number of points awarded

#### Example Request

```sh
curl http://localhost:8080/receipts/{id}/points
```

#### Example Response

```json
{
  "points": 6
}
```

## Running Tests
To run the unit tests, use the following command:
```sh
go test ./internal
```