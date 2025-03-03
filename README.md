# projectUpdates
🔹 Step 1: Owner Login
📌 Endpoint:
http
Copy code
POST /auth/login
📌 Request Body (JSON):
json
Copy code
{
    "email": "owner@example.com",
    "password": "securepassword"
}
📌 Expected Response:
json
Copy code
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
✅ Copy the JWT Token for the next requests.
________________________________________
🔹 Step 2: Create a Library (Owner Only)
📌 Endpoint:
http
Copy code
POST /api/library
📌 Headers:
plaintext
Copy code
Authorization: Bearer <OWNER_TOKEN>
Content-Type: application/json
📌 Request Body (JSON):
json
Copy code
{
    "name": "Central Library"
}
📌 Expected Response:
json
Copy code
{
    "message": "Library created successfully",
    "library": {
        "id": 1,
        "name": "Central Library"
    }
}
✅ Copy the library_id for the next request.
________________________________________
🔹 Step 3: Create an Admin for the Library
📌 Endpoint:
http
Copy code
POST /api/admin
📌 Headers:
plaintext
Copy code
Authorization: Bearer <OWNER_TOKEN>
Content-Type: application/json
📌 Request Body (JSON):
json
Copy code
{
    "email": "admin@example.com",
    "password": "securepassword",
    "role": "admin",
    "library_id": 1
}
📌 Expected Response:
json
Copy code
{
    "message": "Admin registered successfully",
    "admin": {
        "id": 2,
        "email": "admin@example.com",
        "role": "admin",
        "library_id": 1
    }
}
✅ Copy the admin_id for later.
________________________________________
🔹 Step 4: Admin Login
📌 Endpoint:
http
Copy code
POST /auth/login
📌 Request Body (JSON):
json
Copy code
{
    "email": "admin@example.com",
    "password": "securepassword"
}
📌 Expected Response:
json
Copy code
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
✅ Copy the JWT Token for Admin.
________________________________________
🔹 Step 5: Add a Book to the Library (Admin Only)
📌 Endpoint:
http
Copy code
POST /api/book
📌 Headers:
plaintext
Copy code
Authorization: Bearer <ADMIN_TOKEN>
Content-Type: application/json
📌 Request Body (JSON):
json
Copy code
{
    "isbn": "978-3-16-148410-0",
    "title": "Golang Mastery",
    "author": "John Doe",
    "publisher": "Tech Books",
    "copies": 5,
    "library_id": 1
}
📌 Expected Response:
json
Copy code
{
    "message": "Book added successfully",
    "book": {
        "isbn": "978-3-16-148410-0",
        "title": "Golang Mastery",
        "author": "John Doe",
        "publisher": "Tech Books",
        "copies": 5,
        "library_id": 1
    }
}
________________________________________
🔹 Step 6: User Registration
📌 Endpoint:
http
Copy code
POST /api/user
📌 Headers:
plaintext
Copy code
Authorization: Bearer <ADMIN_TOKEN>
Content-Type: application/json
📌 Request Body (JSON):
json
Copy code
{
    "email": "user@example.com",
    "password": "securepassword",
    "role": "user",
    "library_ids": [1]
}
📌 Expected Response:
json
Copy code
{
    "message": "User registered successfully",
    "user": {
        "id": 3,
        "email": "user@example.com",
        "role": "user",
        "library_ids": [1]
    }
}
✅ Copy the user_id.
________________________________________
🔹 Step 7: User Login
📌 Endpoint:
http
Copy code
POST /auth/login
📌 Request Body (JSON):
json
Copy code
{
    "email": "user@example.com",
    "password": "securepassword"
}
📌 Expected Response:
json
Copy code
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
✅ Copy the JWT Token for User.
________________________________________
🔹 Step 8: Search for Books (User Only)
📌 Endpoint:
http
Copy code
GET /api/books/search?title=Golang
📌 Headers:
plaintext
Copy code
Authorization: Bearer <USER_TOKEN>
📌 Expected Response:
json
Copy code
{
    "books": [
        {
            "isbn": "978-3-16-148410-0",
            "title": "Golang Mastery",
            "author": "John Doe",
            "publisher": "Tech Books",
            "copies": 5,
            "library_id": 1
        }
    ]
}
________________________________________
🔹 Step 9: Request to Issue a Book (User Only)
📌 Endpoint:
http
Copy code
POST /api/issue
📌 Headers:
plaintext
Copy code
Authorization: Bearer <USER_TOKEN>
Content-Type: application/json
📌 Request Body (JSON):
json
Copy code
{
    "user_id": 3,
    "isbn": "978-3-16-148410-0",
    "library_id": 1
}
📌 Expected Response:
json
Copy code
{
    "message": "Issue request submitted",
    "request": {
        "user_id": 3,
        "isbn": "978-3-16-148410-0",
        "library_id": 1,
        "request_date": 1700000000
    }
}
________________________________________
🔹 Step 10: Admin Approves the Request
📌 Endpoint:
http
Copy code
PUT /api/issue/approve/:id
📌 Headers:
plaintext
Copy code
Authorization: Bearer <ADMIN_TOKEN>
📌 Expected Response:
json
Copy code
{
    "message": "Issue request approved",
    "issue": {
        "id": 1,
        "user_id": 3,
        "isbn": "978-3-16-148410-0",
        "status": "approved"
    }
}
________________________________________
🔹 Step 11: Admin Issues the Book to User
📌 Endpoint:
http
Copy code
POST /api/issue/book/:isbn
📌 Headers:
plaintext
Copy code
Authorization: Bearer <ADMIN_TOKEN>
📌 Expected Response:
json
Copy code
{
    "message": "Book issued successfully",
    "issue": {
        "id": 1,
        "user_id": 3,
        "isbn": "978-3-16-148410-0",
        "status": "issued"
    }
}

