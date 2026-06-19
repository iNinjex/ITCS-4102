# Programming Languages REST API Comparison Project

## Overview

This project demonstrates how the same RESTful CRUD API can be implemented using three different programming languages and frameworks:

- Python with FastAPI
- JavaScript with Express.js
- Go with Gin

The goal of the project is to compare the syntax, development experience, performance, and tooling of each language while solving the same problem.

---

# Project Objectives

The objectives of this project are to:

- Compare three popular programming languages.
- Demonstrate CRUD (Create, Read, Update, Delete) operations.
- Compare different web frameworks.
- Compare static and dynamic typing.
- Explore package management and development tools.
- Analyze which language is most suitable for different types of applications.

---

# Technologies Used

## Python

- Python 3.11
- FastAPI
- Uvicorn
- Pydantic

## JavaScript

- Node.js
- Express.js
- npm

## Go

- Go
- Gin Framework

---

# Project Structure

```
ProgrammingLanguagesProject
│
├── Docs
├── Proposal
├── Report
├── Screenshots
│
├── Python-FastAPI
├── Javascript-express
└── Go-gin
```

---

# REST API Endpoints

All three implementations provide the same API.

| Method | Endpoint | Description |
|---------|----------|-------------|
| GET | / | Home page |
| GET | /languages | Retrieve all languages |
| GET | /languages/{id} | Retrieve one language |
| POST | /languages | Create a language |
| PUT | /languages/{id} | Update a language |
| DELETE | /languages/{id} | Delete a language |

---

# Running the Python Version

Open a terminal.

```
cd Python-FastAPI
```

Activate the virtual environment.

Windows:

```
.\.venv\Scripts\activate
```

Install dependencies.

```
pip install -r requirements.txt
```

Run the server.

```
python -m uvicorn main:app --reload
```

Open:

```
http://127.0.0.1:8000/docs
```

---

# Running the JavaScript Version

Open a terminal.

```
cd Javascript-express
```

Install dependencies.

```
npm install
```

Run the server.

```
node server.js
```

Open:

```
http://localhost:3000
```

---

# Running the Go Version

Open a terminal.

```
cd Go-gin
```

Run the server.

```
go run main.go
```

Open:

```
http://localhost:8080
```

---

# Testing

The Python implementation was tested using FastAPI Swagger UI.

The JavaScript and Go implementations were tested using Thunder Client in Visual Studio Code.

The following CRUD operations were successfully tested:

- GET
- POST
- PUT
- DELETE

Screenshots of each implementation are included in the Screenshots folder.

---

# Comparison Summary

| Feature | Python | JavaScript | Go |
|----------|--------|------------|----|
| Typing | Dynamic | Dynamic | Static |
| Framework | FastAPI | Express | Gin |
| Performance | Good | Good | Excellent |
| Readability | Excellent | Good | Good |
| Learning Curve | Easy | Moderate | Moderate |
| Best Use | APIs, AI, Data Science | Web Development | High Performance APIs |

---

# Conclusion

Implementing the same REST API in three different programming languages provided insight into how each language approaches web development.

Python offered the simplest and most readable implementation through FastAPI. JavaScript provided flexibility and a large ecosystem through Express. Go offered excellent performance, simple concurrency support, and a clean syntax for backend development.

Although all three implementations achieved the same functionality, each language demonstrated different strengths depending on the application requirements.

---

# Author

Tai Yang

ITCS 4102

Programming Languages

Disclaimer : This project was built with the help of AI.