# JWT-Protected vKYC API in Go

This project is a simple Go-based microservice using **JWT authentication** and the **Gin** web framework. It exposes a secured endpoint for initiating a vKYC (Video Know Your Customer) session. All requests to this endpoint must be authorized using a valid JWT.

---

## 📂 Folder Structure

```
jwt_auth_vkyc/
├── main.go                # Entry point of the application
├── handlers/              # Contains HTTP handlers
│   └── vkyc.go
├── middleware/            # Contains JWT authentication middleware
│   └── auth.go
├── utils/                 # Contains helper function to generate JWT tokens
│   └── jwt.go
```

---

## 🚀 Features

- ✅ JWT-based stateless authentication
- 🔐 Secured endpoint `/v1/initiate-vkyc`
- 🔧 Token generation helper using `utils.GenerateJWT`
- 🛡️ Token validation middleware

---

## 🔑 How JWT Authentication Works

- Clients must include a **valid JWT token** in the `Authorization` header:

```
Authorization: Bearer <eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDc4MDIwNzMsImlhdCI6MTc0Nzc5ODQ3Mywicm9sZSI6ImN1c3RvbWVyIiwic3ViIjoic2VqYWxuYXlhayJ9.Z75V2N6bb3OqRVHdcwvIjJm6gaQqpN-mn8CypKIhteI>
```

- The server verifies the token using a shared HMAC secret.
- If valid, it extracts user info from the token and proceeds.

---

## 🛠️ How to Run This Project

### 1. Clone and Setup
```
git clone <repo_url>
cd jwt_auth_vkyc
go mod tidy
```

### 2. Start the Server
```
go run main.go
```
Server will start at `http://localhost:8080`

### 3. Generate a JWT Token (CLI method)
Create a `generate_token.go`:
```go
package main

import (
  "fmt"
  "jwt_auth_vkyc/utils"
)

func main() {
  token, _ := utils.GenerateJWT("sejalnayak", "customer")
  fmt.Println(token)
}
```
Then run:
```
go run generate_token.go
```

### 4. Make a Secure API Call
Use curl or Postman:

```bash
curl --location 'http://localhost:8080/v1/initiate-vkyc' \
--header 'Authorization: Bearer <PASTE_TOKEN_HERE>' \
--header 'Content-Type: application/json' \
--data '{}'
```

### ✅ Response:
```json
{
  "message": "vKYC session initiated",
  "user_id": "sejalnayak",
  "role": "customer",
  "vkyc_link": "https://vkyc.nayak.com/v1/video?token=securetoken#step1"
}
```

---

## 🧪 Optional: Add Token Generator Endpoint
To make development easier, you can add this route in `main.go`:
```go
r.GET("/generate-token", func(c *gin.Context) {
  token, _ := utils.GenerateJWT("sejalnayak", "customer")
  c.JSON(200, gin.H{"token": token})
})
```

Then hit `http://localhost:8080/generate-token` in your browser to get a token.

---

## 📄 License
This is a practice/demo project for understanding JWT authentication in Go. No external storage or encryption is implemented.
# jwt_auth_vkyc
