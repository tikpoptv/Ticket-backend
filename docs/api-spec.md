# API Specification

## Base URL
```
http://localhost:8080
```

## Authentication
Currently, the API is open for registration without authentication.

## Endpoints

### 1. Welcome Message
```
GET /
Access: Public
Response: { "message": "Welcome to the API!" }
```

### 2. Register
```
POST /register
Access: Public

Request:
{
    "username": "string",     // min: 3, max: 64
    "password": "string",     // min: 6
    "name": "string",
    "email": "string"        // valid email
}

Response:
201: { "message": "registration successful" }
400: { "error": "error message" }
```

Possible Error Messages:
- "username already exists"
- "invalid request body"
- "validation error" 