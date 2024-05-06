# API Endpoints

## `GET /`

Returns a simple greeting message.

**Response:**

```json
"Hello, World!"
```

## `POST /login`

Authenticates a user and returns a JWT token.

**Request Body:**

```json
{
    "email": "body@gmail.com",
    "password": "123456"
}
```

**Response:**

```json
{
    "token": "jwt.token.here",
    "user": {
        "id": 1,
        "email": "user@example.com"
    }
}
```

## `POST /hello`

Returns a protected greeting message. Requires a valid JWT token in the `Authorization` header.

**Headers:**

```
Authorization: Bearer jwt.token.here
```



**

Response:**

```json
"Hello, Protect!"
```
