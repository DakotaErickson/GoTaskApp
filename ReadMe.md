# Go REST Backend for Task App

This project provides a basic RESTful API for managing tasks using Go and MongoDB.

## Getting Started

1. **Prerequisites:**
   - Go installed on your system.
   - MongoDB installed and running.
   - `air` command-line tool installed for development (optional, but recommended).

2. **Clone the Repository:**

3. **Install Dependencies:**

    ```go mod tidy```

4. **Create a `.env` file**
    - MONGODB_URI=`<Your connection string>`
    - PORT=`<Whatever port you want it to run on>`

5. **Run the application**
    - Using air: `air`
    - Without air: `go run main.go`

6. **Interact with the API**

Example Endpoints:

    GET /api/todos: Get a list of all todos.
    POST /api/todos: Create a new todo.
    PATCH /api/todos/:id: Mark an existing todo as complete.
    DELETE /api/todos/:id: Delete a todo.