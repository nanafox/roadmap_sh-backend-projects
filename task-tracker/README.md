# Task CLI - A roadmap.sh Backend Project

**[Project Link](https://roadmap.sh/projects/task-tracker)**

Task CLI is a simple command-line application for managing tasks. It allows
users to add, list, update, mark as in-progress, mark as done, and delete
tasks. The application saves tasks to a persistent storage file for reuse
across sessions.

---

## Features

- **Add Task**: Create a new task with a description.
- **List Tasks**: View all tasks, including their statuses.
- **Update Task**: Update the description of an existing task.
- **Mark Task**: Change the status of a task to "in-progress" or "done."
- **Delete Task**: Remove a task permanently.

---

## Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/nanafox/roadmap_sh-backend-projects.git
   cd roadmap_sh-backend-projects/task-tracker
   ```

2. **Build the Application**  
   Make sure you have Go installed on your machine. Then build the application:

   ```bash
   go build -o task-cli
   ```

3. **Run the Application**  
   Execute commands using the built binary:
   ```bash
   ./task-cli <command>
   ```

---

## Usage

All commands follow the syntax:

```bash
task-cli <action> <arguments...>
```

### Commands

#### Add a Task

Create a new task with a description.

```bash
task-cli add <description>
```

- Example:
  ```bash
  task-cli add "Buy groceries"
  ```

#### List All Tasks

View all tasks and their statuses.

```bash
task-cli list
```

- Example:
  ```bash
  task-cli list
  ```

#### Update a Task

Update the description of an existing task.

```bash
task-cli update <taskId> <new description>
```

- Example:
  ```bash
  task-cli update 1 "Buy groceries and vegetables"
  ```

#### Mark a Task as In-Progress

Change the status of a task to "in-progress."

```bash
task-cli mark-in-progress <taskId>
```

- Example:
  ```bash
  task-cli mark-in-progress 1
  ```

#### Mark a Task as Done

Change the status of a task to "done."

```bash
task-cli mark-done <taskId>
```

- Example:
  ```bash
  task-cli mark-done 1
  ```

#### Delete a Task

Remove a task permanently from storage.

```bash
task-cli delete <taskId>
```

- Example:
  ```bash
  task-cli delete 1
  ```

---

## How It Works

### Storage

Tasks are persisted to a JSON file (`tasks.json`) in the project directory. The application reads from and writes to this file during each operation.

### Task Structure

Each task is stored with the following fields:

- **Id**: Unique identifier for the task.
- **Description**: Description of the task.
- **Status**: Current status of the task ("pending", "in-progress", or "done").
- **CreatedAt**: Timestamp when the task was created.
- **UpdatedAt**: Timestamp when the task was last updated.

---

## Adding the Delete Functionality

The `delete` command will allow users to remove a task permanently. Below is a description of the function you will implement:

### Function: `DeleteTask`

#### Usage:

```bash
task-cli delete <taskId>
```

#### Implementation Overview:

- Validate that `cmd` contains the correct number of arguments.
- Parse the `taskId` from the command arguments.
- Attempt to find the task by ID.
- If the task exists, delete it from the storage.
- Save the updated task list back to the JSON file.

#### Example:

- Input:
  ```bash
  task-cli delete 2
  ```
- Output:
  ```
  Task with ID 2 deleted successfully!
  ```

---

## Error Handling

- **Invalid Input**: Returns a descriptive error for malformed commands.
- **Task Not Found**: Returns an error if the specified task does not exist.
- **File Errors**: Handles issues with reading or writing the storage file.

---

## Future Improvements

- Add support for task priorities.
- Enable filtering tasks by status (e.g., "pending", "done").
- Enhance storage format (e.g., use a database instead of a JSON file).
- Provide colored output for better readability in the terminal.

---

## Contributing

Contributions are welcome! To contribute:

1. Fork the repository.
2. Create a feature branch:
   ```bash
   git checkout -b feature-name
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add new feature"
   ```
4. Push to the branch:
   ```bash
   git push origin feature-name
   ```
5. Open a pull request.

---

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
