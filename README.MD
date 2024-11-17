# Order Management System - API Documentation
This document provides an overview of the Order Management System API endpoints and instructions for making requests.

## API Endpoints

- The API is accessible at {{your_server_address}}/api/v1 (replace with your actual server address if deployed elsewhere).

## Authentication

- Register a new user: POST /auth/signup (Requires JSON body with email and password)
- Login and obtain an authorization token: POST /auth/signin (Requires JSON body with email and password)

## Authorization

- Most Order Management API endpoints require authorization. Include the obtained token in the Authorization header following the format Bearer <token>.

## Orders

- Create a new order: POST /orders (Requires authorization and JSON body with order details)
- Retrieve a list of orders: GET /orders/all?limit=10&page=1 (Optional query parameters for pagination)
- Cancel an existing order: PUT /orders/:consignment_id/cancel (Requires authorization and replaces :consignment_id with the actual order ID)

## Example Usage (using curl)

### Register a new user

```bash
curl --location 'http://localhost:8081/api/v1/auth/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
  "email":"skza45@gmail.com",
  "password": "123456"
}'
```
### Example Response 

#### 200 OK
```json
{
    "code": 201,
    "message": "User created successfully"
}
```

#### 400 Bad Request
```json
{
    "code": 400,
    "error": "Email or password is invalid"
}
```


### Login and obtain authorization token

```bash
curl --location 'http://localhost:8081/api/v1/auth/signin' \
--header 'Content-Type: application/json' \
--data-raw '{
  "email":"skza45@gmail.com",
  "password": "123456"
}'
```

### Example Response        

#### 200 OK

```json
{
    "code": 200,
    "message": "Successfully signed in",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzE5NTYxNDksInN1YiI6MX0.fW2cLrfMXonaipC2BbgodX4D212K6qAXuktQrDfZLyk"
}
```

#### 400 Bad Request

```json
{
    "code": 400,
    "error": "Email or password is invalid"
}
```

### Create a new order

```bash
curl --location 'http://localhost:8081/api/v1/orders/' \
--header 'Authorization: Bearer YOUR_TOKEN' \
--header 'Content-Type: application/json' \
--data '{
  "recipient_name": "John Doe",
  "recipient_phone": "01783092354",
  "recipient_address": "123 Main St, chittagong",
  "special_instruction": "Leave at the door",
  "amount_to_collect": 100,
  "item_description": "Books",
  "merchant_order_id":"MStSDHJK"
}'
```

### Example Response

#### 200 OK

```json
{
    "code": 200,
    "data": {
        "consignment_id": "a91c7c0e224f",
        "delivery_fee": 100,
        "merchant_order_id": "MStSDHJK",
        "order_status": "Pending"
    },
    "message": "Order Created Successfully",
    "type": "success"
}
```

#### 422 Unprocessable Entity

```json
{
    "code": 422,
    "errors": {
        "recipient_phone": [
            "Invalid phone number format."
        ]
    },
    "message": "Please fix the given errors",
    "type": "error"
}
```

#### 401 Unauthorized

```json
{
    "error": "Unauthorized. Token is missing"
}
```

### Get a list of orders

```bash
curl --location 'http://localhost:8081/api/v1/orders/all?limit=10&page=1' \
--header 'Authorization: Bearer YOUR_TOKEN'
```

### Example Response

#### 200 OK

```json
{
    "code": 200,
    "data": {
        "current_page": 2,
        "data": [
            {
                "archived": false,
                "cod_fee": 1,
                "delivery_fee": 100,
                "instruction": "Leave at the door",
                "item_type": 2,
                "merchant_order_id": "MStSDHJK",
                "order_amount": 100,
                "order_consignment_id": "85fa2c0e40b7",
                "order_created_at": "2024-11-17 22:05:56.355705+08",
                "order_description": "Books",
                "order_status": "Canceled",
                "order_type": 48,
                "order_type_id": 48,
                "promo_discount": 0,
                "recipient_address": "123 Main St, chittagong",
                "recipient_name": "John Doe",
                "recipient_phone": "01783092354",
                "total_fee": 199
            },
            {
                "archived": false,
                "cod_fee": 1,
                "delivery_fee": 100,
                "instruction": "Leave at the door",
                "item_type": 2,
                "merchant_order_id": "MStSDHJK",
                "order_amount": 100,
                "order_consignment_id": "a91c7c0e224f",
                "order_created_at": "181812-52-1812 1112:52:1212",
                "order_description": "Books",
                "order_status": "Pending",
                "order_type": 48,
                "order_type_id": 48,
                "promo_discount": 0,
                "recipient_address": "123 Main St, chittagong",
                "recipient_name": "John Doe",
                "recipient_phone": "01783092354",
                "total_fee": 199
            }
        ],
        "last_page": 2,
        "per_page": 3,
        "total": 5,
        "total_in_page": 2
    },
    "message": "Orders successfully fetched.",
    "type": "success"
}   
```


### Cancel an order

```bash
curl --location --request PUT 'http://localhost:8081/api/v1/orders/85fa2c0e407/cancel' \
--header 'Authorization: Bearer YOUR_TOKEN'
```

### Example Response

#### 400 OK

```json
{
    "code": 400,
    "message": "Please contact cx to cancel order",
    "type": "error"
}
```

### 200 OK

```json
{
    "code": 200,
    "message": "Order Cancelled Successfully",
    "type": "success"
}
``` 
        