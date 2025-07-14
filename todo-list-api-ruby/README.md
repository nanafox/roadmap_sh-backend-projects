# Todo List API

A RESTful API for managing todo tasks built with Ruby on Rails. This project includes user authentication using Rodauth and allows users to create, read, update, and delete tasks with different priorities and statuses.

This is the solution for [roadmap.sh Todo List API](https://roadmap.sh/projects/todo-list-api) challenge.

## Features

- User authentication (register, login, logout, password reset)
- CRUD operations for tasks
- Task priorities (low, medium, high)
- Task statuses (pending, in-progress, completed)
- JSON API responses
- User-specific task management

## Requirements

- Ruby 3.4.4
- Rails 8.0.2
- SQLite3 database

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/nanafox/roadmap_sh-backend-projects.git
   cd roadmap_sh-backend-projects/todo-list-api-ruby
   ```

2. Install dependencies:

   ```bash
   bundle install
   ```

3. Setup the database:

   ```bash
   rails db:create
   rails db:migrate
   rails db:seed
   ```

4. Start the server:

   ```bash
   rails server
   ```

The API will be available at `http://localhost:3000`

## API Endpoints

### Authentication

- `POST /auth/login` - Login user
- `POST /auth/create-account` - Register new user
- `POST /auth/logout` - Logout user
- `POST /auth/reset-password` - Reset password

### Tasks

- `GET /tasks` - Get all tasks for authenticated user
- `GET /tasks/:id` - Get specific task
- `POST /tasks` - Create new task
- `PUT /tasks/:id` - Update task
- `DELETE /tasks/:id` - Delete task

### Task Parameters

- `name` (required) - Task name
- `priority` (required) - Task priority (low, medium, high)
- `status` (required) - Task status (pending, in-progress, completed)

## Example Usage

### Create a task

```bash
curl -X POST http://localhost:3000/tasks \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d '{
    "name": "Complete project",
    "priority": "high",
    "status": "pending",
  }'
```

### Get all tasks

```bash
curl -X GET http://localhost:3000/tasks \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## Running Tests

```bash
rails test
```
