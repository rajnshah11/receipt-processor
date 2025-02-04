Here’s the updated `README.md` file for your Receipt Processor project:

---

# **Receipt Processor**

## **Overview**

The Receipt Processor is a web service built in Go that processes receipts and calculates points based on specific rules. It provides two API endpoints:

- **POST /receipts/process**: Submits a receipt for processing and returns a unique ID.
- **GET /receipts/{id}/points**: Retrieves the points awarded for a specific receipt using its unique ID.

The application is Dockerized for easy setup and deployment.

---

## **Features**

- Processes receipts and calculates points based on specific rules.
- In-memory storage of processed receipts (data is lost when the server restarts).
- Fully Dockerized for portability.
- Simple RESTful API with JSON input/output.

---

## **Rules for Calculating Points**

1. **Retailer Name**: 1 point for every alphanumeric character in the retailer name.
2. **Round Dollar Total**: 50 points if the total is a round dollar amount with no cents.
3. **Multiple of $0.25**: 25 points if the total is a multiple of $0.25.
4. **Item Count**: 5 points for every two items on the receipt.
5. **Item Description Length**: If the trimmed length of an item description is a multiple of 3, multiply the item price by 0.2 and round up to the nearest integer. Add this to the points.
6. **Odd Purchase Day**: 6 points if the day of purchase is odd.
7. **Purchase Time**: 10 points if the time of purchase is between 2:00 PM and 4:00 PM.

---

## **Project Structure**

```
receipt-processor/
├── main.go               # Main Go application code
├── controllers/          # Contains HTTP handler logic
│   └── receipt_controller.go
├── models/               # Contains data models
│   └── receipt.go
├── services/             # Contains business logic and receipt processing
│   └── receipt_service.go
├── go.mod                # Go module file
├── go.sum                # Dependency checksums
├── Dockerfile            # Docker configuration
└── README.md             # Documentation (this file)
```

---

## **Prerequisites**

- Go (version 1.20 or higher)
- Docker

---

## **Setup Instructions**

### **1. Clone the Repository**
```bash
git clone <repository-url>
cd receipt-processor
```

### **2. Build and Run Locally (Optional)**

If you want to run the application locally without Docker:

#### Install dependencies:
```bash
go mod tidy
```

#### Run the application:
```bash
go run main.go
```

The server will start at `http://localhost:8000`.

---

### **3. Build and Run with Docker**

#### Build the Docker image:
```bash
docker build -t receipt-processor .
```

#### Run the container:
```bash
docker run -p 8000:8000 receipt-processor
```

The application will be accessible at `http://localhost:8000`.

---

## **API Endpoints**

### **1. POST /receipts/process**

Submits a receipt for processing.

#### Request:
- Header: `Content-Type: application/json`
- Body:
```json
{
    "retailer": "Target",
    "purchaseDate": "2022-01-01",
    "purchaseTime": "13:01",
    "items": [
        {"shortDescription": "Mountain Dew", "price": "6.49"},
        {"shortDescription": "Pizza", "price": "12.25"}
    ],
    "total": "18.74"
}
```

#### Response:
```json
{ "id": "<generated-id>" }
```

---

### **2. GET /receipts/{id}/points**

Retrieves the points awarded for a specific receipt.

#### Request:
- URL Path Parameter: `{id}` (the ID returned from `/receipts/process`).

#### Example:
```
GET http://localhost:8000/receipts/<generated-id>/points
```

#### Response:
```json
{ "points": <calculated-points> }
```

---

## **Testing**

### Using `curl`:

#### Submit a receipt:
```bash
curl -X POST -H "Content-Type: application/json" \
-d '{
    "retailer": "Target",
    "purchaseDate": "2022-01-01",
    "purchaseTime": "13:01",
    "items": [
        {"shortDescription": "Mountain Dew", "price": "6.49"},
        {"shortDescription": "Pizza", "price": "12.25"}
    ],
    "total": "18.74"
}' \
http://localhost:8000/receipts/process
```

#### Retrieve points using the returned ID:
```bash
curl http://localhost:8000/receipts/<generated-id>/points
```

---

### Using Postman:

1. Create a new POST request to `http://localhost:8000/receipts/process`.
   - Add a JSON body as shown above.
   - Send the request and copy the returned `<generated-id>`.

2. Create a new GET request to `http://localhost:8000/receipts/<generated-id>/points`.
   - Send the request to retrieve points.

---

## **Common Issues**

### Docker Build Fails (`go.mod` Error):
- Ensure your `go.mod` file has a valid Go version (e.g., `go 1.20`).
- Run `go mod tidy` locally to fix dependency issues.

### Port Already in Use:
If port `8000` is already in use, use another port when running Docker:
```bash
docker run -p 9000:8000 receipt-processor
```
Access it at `http://localhost:9000`.

### Data Loss After Restart:
The application uses in-memory storage, so data is lost when the server restarts.
To persist data, consider integrating a database (e.g., SQLite or Redis).

---

## **Future Enhancements**

1. Add persistent storage (e.g., SQLite or Redis) to retain data across server restarts.
2. Implement authentication and authorization for secure access.
3. Add more robust error handling and validation for inputs.

---

## **License**

This project is licensed under the MIT License.

---

This README provides all necessary information to set up, run, and test your Receipt Processor project! Let me know if you need further assistance!