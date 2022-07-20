# Middleware & Books CRUD

## Rest API Documentation

In this service there are 1 endpoint without authentication and 4 endpoints with authentication

#### A. GET All Users

- **Endpoint:** /users
- **Method:** GET
- **Example Request:** http://localhost:8000/users
- **Example Success Response (Status Code: 200**)
```json
{
    "message": "success get all data",
    "data": [
        {
            "id": 1,
            "name": "Wahyu",
            "email": "wahyuagung26@gmail.com",
            "roles": {
                "books": {
                    "create": true,
                    "delete": true,
                    "update": true,
                    "view": true
                }
            },
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZW1haWwiOiJ3YWh5dWFndW5nMjZAZ21haWwuY29tIiwicm9sZXMiOiJ7XCJib29rc1wiOntcImNyZWF0ZVwiOnRydWUsXCJ1cGRhdGVcIjp0cnVlLFwiZGVsZXRlXCI6dHJ1ZSxcInZpZXdcIjp0cnVlfX0iLCJleHAiOjE2NTg1ODY0NjN9.yhr8D8NKAIe2qQ1NuA8N7l36LYLoaDI2ZsIZbh6hopg"
        },
        {
            "id": 2,
            "name": "Agung",
            "email": "wahyu.agung@majoo.id",
            "roles": {
                "books": {
                    "create": false,
                    "delete": false,
                    "update": true,
                    "view": true
                }
            },
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiZW1haWwiOiJ3YWh5dS5hZ3VuZ0BtYWpvby5pZCIsInJvbGVzIjoie1wiYm9va3NcIjp7XCJjcmVhdGVcIjpmYWxzZSxcInVwZGF0ZVwiOnRydWUsXCJkZWxldGVcIjpmYWxzZSxcInZpZXdcIjp0cnVlfX0iLCJleHAiOjE2NTg1ODY0NjN9.__3QWNfZ5X1h6RyslHVRwrSxTOplGUUmgtPBsiuvI2w"
        }
    ]
}
```
#### B. GET All Books

- **Endpoint:** /books
- **Method:** GET
- **Header Reqest:** Authorization: Bearer {token}
- **Example Request:** http://localhost:8000/books
- **Example Success Response (Status Code: 200**)
```json
{
    "message": "success get all data",
    "data": [
        {
            "id": 1,
            "title": "Peekaboo Whats in the Snow",
            "price": 46800,
            "author": "Tim Pelangi Indonesia",
            "publisher": "PELANGI INDONESIA",
            "publish_date": "2022-03-01T00:00:00+07:00",
            "created_at": "2022-07-20T22:12:46.22+07:00",
            "updated_at": "2022-07-20T22:12:46.22+07:00"
        },
        {
            "id": 2,
            "title": "Pengantar Ilmu Tafsir",
            "price": 43250,
            "author": "Drs. A. Fudlali",
            "publisher": "Angkasa",
            "publish_date": "2005-01-01T00:00:00+07:00",
            "created_at": "2022-07-20T22:12:46.22+07:00",
            "updated_at": "2022-07-20T22:12:46.22+07:00"
        },
        {
            "id": 3,
            "title": "The Miracle Of Ikhlas",
            "price": 29325,
            "author": "Anin DP",
            "publisher": "Mueeza",
            "publish_date": "2021-02-01T00:00:00+07:00",
            "created_at": "2022-07-20T22:12:46.22+07:00",
            "updated_at": "2022-07-20T22:12:46.22+07:00"
        }
    ]
}
```
- **Example Error Response (Status Code: 401**)
```json
{
    "message": "invalid or expired jwt",
}
```
- **Example Error Response (Status Code: 403**)
```json
{
    "message": "You don't have credential to access this page",
}
```
#### C. GET Spesific Books

- **Endpoint:** /books/1
- **Method:** GET
- **Header Reqest:** Authorization: Bearer {token}
- **Example Request:** http://localhost:8000/books/1
- **Example Success Response (Status Code: 200**)
```json
{
    "message": "success get all data",
    "data": [
        {
            "id": 1,
            "title": "Peekaboo Whats in the Snow",
            "price": 46800,
            "author": "Tim Pelangi Indonesia",
            "publisher": "PELANGI INDONESIA",
            "publish_date": "2022-03-01T00:00:00+07:00",
            "created_at": "2022-07-20T22:12:46.22+07:00",
            "updated_at": "2022-07-20T22:12:46.22+07:00"
        }
    ]
}
```
- **Example Error Response (Status Code: 401**)
```json
{
    "message": "invalid or expired jwt",
}
```
- **Example Error Response (Status Code: 403**)
```json
{
    "message": "You don't have credential to access this page",
}
```
- **Example Error Response (Status Code: 404**)
```json
{
    "message": "book not found",
}
```
#### D. Create New Books

- **Endpoint:** /books
- **Method:** POST
- **Header Reqest:** Authorization: Bearer {token}
- **Content-Type:** application/json
- **Example Request:** http://localhost:8000/books
- **Example Raw Body:**
```json
{
    "title": "Peekabo22o Whats in the 12",
    "price": 46800,
    "author": "Tim Pelangi Indonesia",
    "publisher": "PELANGI INDONESIA",
    "publish_date": "2022-03-01T00:00:00+07:00"
}
```
- **Example Success Response (Status Code: 200**)
```json
{
    "message": "success create new book",
    "data": {
        "id": 4,
        "title": "Peekabo22o Whats in the 12",
        "price": 46800,
        "author": "Tim Pelangi Indonesia",
        "publisher": "PELANGI INDONESIA",
        "publish_date": "2022-03-01T00:00:00+07:00",
        "created_at": "2022-07-20T22:34:03.321+07:00",
        "updated_at": "2022-07-20T22:34:03.321+07:00"
    }
}
```
- **Example Error Response (Status Code: 400**)
```json
{
    "message": "Key: 'Books.Price' Error:Field validation for 'Price' failed on the 'required' tag",
    "data": {}
}
```
- **Example Error Response (Status Code: 401**)
```json
{
    "message": "invalid or expired jwt",
}
```
- **Example Error Response (Status Code: 403**)
```json
{
    "message": "You don't have credential to access this page",
}
```
#### E. Update Existing Books

- **Endpoint:** /books
- **Method:** PUT
- **Header Reqest:** Authorization: Bearer {token}
- **Content-Type:** application/json
- **Example Request:** http://localhost:8000/books
- **Example Raw Body:**
```json
{
    "id" : 1,
    "title": "Peekabo22o Whats in the 12",
    "price": 500000,
    "author": "Tim Pelangi Indonesia",
    "publisher": "PELANGI INDONESIA",
    "publish_date": "2022-03-01T00:00:00+07:00"
}
```
- **Example Success Response (Status Code: 200**)
```json
{
    "message": "success update book",
    "data": {
        "id": 1,
        "title": "Peekabo22o Whats in the 12",
        "price": 500000,
        "author": "Tim Pelangi Indonesia",
        "publisher": "PELANGI INDONESIA",
        "publish_date": "2022-03-01T00:00:00+07:00",
        "created_at": "2022-07-20T22:34:03.321+07:00",
        "updated_at": "2022-07-20T22:34:03.321+07:00"
    }
}
```
- **Example Error Response (Status Code: 400**)
```json
{
    "message": "Key: 'Books.Price' Error:Field validation for 'Price' failed on the 'required' tag",
    "data": {}
}
```
- **Example Error Response (Status Code: 401**)
```json
{
    "message": "invalid or expired jwt",
}
```
- **Example Error Response (Status Code: 403**)
```json
{
    "message": "You don't have credential to access this page",
}
```
#### E. Delete Existing Books

- **Endpoint:** /books/{:id}
- **Method:** DELETE
- **Header Reqest:** Authorization: Bearer {token}
- **Content-Type:** application/json
- **Example Request:** http://localhost:8000/books/1
- **Example Success Response (Status Code: 200**)
```json
{
    "message": "success to delete book",
    "data": true
}
```
- **Example Error Response (Status Code: 401**)
```json
{
    "message": "invalid or expired jwt",
}
```
- **Example Error Response (Status Code: 403**)
```json
{
    "message": "You don't have credential to access this page",
}
```