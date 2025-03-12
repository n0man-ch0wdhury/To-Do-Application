# Todo Application

A full-stack Todo application built with Go (backend) and React (frontend). This application allows users to register, login, and manage their todo items with CRUD operations.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
  - [Backend Setup](#backend-setup)
  - [Frontend Setup](#frontend-setup)
- [API Documentation](#api-documentation)
- [Authentication Flow](#authentication-flow)
- [Usage](#usage)

## Features

- User authentication (Register, Login, Logout)
- Create, Read, Update, and Delete todo items
- Mark todos as completed
- Responsive UI built with Material-UI
- JWT-based authentication
- RESTful API

## Tech Stack

### Backend
- Go (1.24)
- Gorilla Mux (Router)
- PostgreSQL (Database)
- JWT for authentication

### Frontend
- React 18
- React Router for navigation
- Material-UI for UI components
- Axios for API requests
- Vite as build tool

## Project Structure

```
├── .env                  # Environment variables
├── controllers/          # API controllers
├── database/             # Database connection and operations
├── frontend/             # React frontend application
├── middleware/           # Authentication middleware
├── models/               # Data models
├── repository/           # Data access layer
├── go.mod                # Go module definition
├── go.sum                # Go module checksums
├── main.go               # Main application entry point
```

## Prerequisites

- Go 1.24 or higher
- Node.js 16 or higher
- npm or yarn
- PostgreSQL database

## Installation

### Backend Setup

1. Clone the repository:

```bash
git clone <repository-url>
cd to-do-applicatoin
```

2. Create a `.env` file in the root directory with the following variables:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=todo_db
SERVER_PORT=8080
JWT_SECRET=your_jwt_secret
```

3. Install Go dependencies:

```bash
go mod download
```

4. Set up the PostgreSQL database:

```sql
CREATE DATABASE todo_db;
```

5. Run the backend server:

```bash
go run main.go
```

The server will start on http://localhost:8080

### Frontend Setup

1. Navigate to the frontend directory:

```bash
cd frontend
```

2. Install dependencies:

```bash
npm install
# or
yarn install
```

3. Start the development server:

```bash
npm run dev
# or
yarn dev
```

The frontend development server will start on http://localhost:5173

## API Documentation

### Authentication Endpoints

#### Register a new user
- **URL**: `/api/auth/register`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "username": "example",
    "email": "example@example.com",
    "password": "password123"
  }
  ```
- **Response**: User details with JWT token

#### Login
- **URL**: `/api/auth/login`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "email": "example@example.com",
    "password": "password123"
  }
  ```
- **Response**: User details with JWT token

#### Logout
- **URL**: `/api/auth/logout`
- **Method**: `POST`
- **Headers**: `Authorization: Bearer <token>`
- **Response**: Success message

### Todo Endpoints

#### Create a new todo
- **URL**: `/api/todos`
- **Method**: `POST`
- **Headers**: `Authorization: Bearer <token>`
- **Request Body**:
  ```json
  {
    "title": "Task title",
    "description": "Task description"
  }
  ```
- **Response**: Created todo item

#### Get all todos
- **URL**: `/api/todos`
- **Method**: `GET`
- **Headers**: `Authorization: Bearer <token>`
- **Response**: Array of todo items

#### Get a specific todo
- **URL**: `/api/todos/{id}`
- **Method**: `GET`
- **Headers**: `Authorization: Bearer <token>`
- **Response**: Todo item details

#### Update a todo
- **URL**: `/api/todos/{id}`
- **Method**: `PUT`
- **Headers**: `Authorization: Bearer <token>`
- **Request Body**:
  ```json
  {
    "title": "Updated title",
    "description": "Updated description",
    "completed": true
  }
  ```
- **Response**: Updated todo item

#### Delete a todo
- **URL**: `/api/todos/{id}`
- **Method**: `DELETE`
- **Headers**: `Authorization: Bearer <token>`
- **Response**: Success message

## Authentication Flow

1. **Registration**: User registers with username, email, and password
2. **Login**: User logs in with email and password, receives JWT token
3. **API Requests**: JWT token is included in the Authorization header for protected routes
4. **Token Validation**: Server validates the token for each protected request
5. **Logout**: Token is invalidated on the server

## Usage

### User Registration and Login

1. Navigate to the application URL
2. Click on "Register" to create a new account
3. Fill in the registration form with your details
4. After registration, you'll be redirected to the login page
5. Enter your credentials to log in

### Managing Todos

1. After logging in, you'll see your todo list (if any)
2. To add a new todo, fill in the "Add New Todo" form and click "Add Todo"
3. To mark a todo as completed, click the checkbox next to it
4. To edit a todo, click the edit icon, make your changes, and click "Save"
5. To delete a todo, click the delete icon

### Logging Out

Click the "Logout" button in the header to log out of the application.

## License

[MIT](LICENSE)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.