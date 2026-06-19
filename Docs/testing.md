# Testing Documentation

## Overview

Each implementation of the Programming Languages REST API was tested to verify that all CRUD (Create, Read, Update, Delete) operations functioned correctly.

The same endpoints were implemented and tested in all three programming languages.

---

# Python (FastAPI)

### Testing Tool

FastAPI Swagger UI

URL:

```
http://127.0.0.1:8000/docs
```

### Tested Endpoints

- GET /
- GET /languages
- GET /languages/{id}
- POST /languages
- PUT /languages/{id}
- DELETE /languages/{id}

### Result

All endpoints returned the expected responses.

Screenshots are located in:

```
Screenshots/Python
```

---

# JavaScript (Express)

### Testing Tool

Thunder Client (Visual Studio Code)

Server:

```
http://localhost:3000
```

### Tested Endpoints

- GET /
- GET /languages
- GET /languages/:id
- POST /languages
- PUT /languages/:id
- DELETE /languages/:id

### Result

All CRUD operations completed successfully.

Screenshots are located in:

```
Screenshots/Javascript
```

---

# Go (Gin)

### Testing Tool

Thunder Client (Visual Studio Code)

Server:

```
http://localhost:8080
```

### Tested Endpoints

- GET /
- GET /languages
- GET /languages/:id
- POST /languages
- PUT /languages/:id
- DELETE /languages/:id

### Result

All CRUD operations completed successfully.

Screenshots are located in:

```
Screenshots/Go
```

---

# Testing Summary

Every implementation successfully performed the following operations:

- Retrieve all programming languages
- Retrieve a single programming language
- Create a new programming language
- Update an existing programming language
- Delete a programming language

No functional issues were encountered after implementation and debugging.

The screenshots included with this project provide evidence that each endpoint was tested successfully.