# Expense Tracker CLI - A roadmap.sh Backend Project

<!--toc:start-->

## Table of Contents

- [Expense Tracker CLI - A roadmap.sh Backend Project](#table-of-contents)
  - [Features](#features)
  - [Installation](#installation)
  - [Usage](#usage)
    - [`add`](#add)
    - [`list`](#list)
    - [`delete`](#delete)
    - [`update`](#update)
    - [`summary`](#summary)
  - [Data Storage](#data-storage)
  - [Error Handling](#error-handling)
  - [Contributing](#contributing)
  - [License](#license)
  <!--toc:end-->

**Expense Tracker** is a simple Command-Line Interface (CLI) application
for managing and tracking your expenses. The tool allows users to easily add,
view, update, delete, and summarize their expenses from the terminal.

The project
**[Expense Tracker](https://roadmap.sh/projects/expense-tracker)**
is a challenge from [roadmap.sh](https://roadmap.sh)

## Features

1. **Add Expenses**
   Add a new expense with a description and amount.

2. **List Expenses**
   View all the tracked expenses in a tabular format.

3. **Delete Expenses**
   Remove an expense by its ID. _Caution: Deletion is irreversible._

4. **Update Expenses**
   Modify an existing expense, useful for correcting mistakes.

5. **Summary**
   View the total amount spent across all tracked expenses or for a specific month.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/nanafox/roadmap_sh-backend-projects.git
   cd expense-tracker
   ```

2. Build the binary:

   ```bash
   go build -o expense-tracker
   ```

3. Run the application:

   ```bash
   ./expense-tracker
   ```

## Usage

The CLI supports the following commands. Use `--help` with any command for
additional details.

### `add`

Adds a new expense to the list of tracked expenses. Requires a description
and an amount.
**Usage:**

```bash
./expense-tracker add --description <description> --amount <amount>
```

**Example:**

```bash
./expense-tracker add --description "Groceries" --amount 50.75
```

### `list`

Lists all the tracked expenses in a tabular format.
**Usage:**

```bash
./expense-tracker list
```

### `delete`

Deletes an expense by its ID. _Caution: This operation is irreversible._
**Usage:**

```bash
./expense-tracker delete --id <expense_id>
```

**Example:**

```bash
./expense-tracker delete --id 3
```

### `update`

Updates an existing expense, allowing you to correct mistakes in the
description or amount.
**Usage:**

```bash
./expense-tracker update --id <expense_id> --description <description> --amount <amount>
```

**Example:**

```bash
./expense-tracker update --id 5 --description "Rent Payment" --amount 750
```

### `summary`

Calculates and prints the total amount spent. Optionally, summarize
expenses for a specific month **in the current year.**
**Usage:**

```bash
./expense-tracker summary [--month <month_number>]
```

**Examples:**

```bash
# Print the total expenses
./expense-tracker summary

# Print the expenses for August (month 8)
./expense-tracker summary --month 8
```

## Data Storage

The application saves all expenses in a JSON file located in the user’s home
directory (`~/.expenses.json`). The data structure ensures fast lookups and
persistence between sessions.

## Error Handling

1. Invalid commands or missing arguments will result in descriptive error messages.
2. Operations such as `delete` or `update` on non-existent IDs will notify you
   that the expense does not exist.

## Contributing

1. Fork the repository.
2. Create a new branch for your feature: `git checkout -b feature-name`
3. Commit your changes: `git commit -m "Add feature-name"`
4. Push to your branch: `git push origin feature-name`
5. Create a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

---

This README provides all the necessary information for using and contributing
to the **Expense Tracker CLI**. For more details, consult the `--help` options
available within the app. Happy expense tracking!
