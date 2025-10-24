# TaskList CLI

A command-line task management application built with Go and Cobra.

## Features

- Create, list, update and delete tasks
- Mark tasks as completed 
- Persistent storage using JSON
- Simple and intuitive CLI interface

## Installation

Requires Go 1.24.5 or higher.

```sh
go get github.com/username/tasklist-cli
```

## Usage

### Add a task

```sh
tasklist-cli add "Buy groceries" 
```

### List all tasks

```sh
tasklist-cli list
```

Output format:
```
[ ] 1: Buy groceries (Created: 2025-10-24 10:55:57, Updated: 2025-10-24 10:55:57)
[x] 2: Conquer Austria (Created: 2025-10-24 11:53:59, Updated: 2025-10-24 11:54:19)
```

### Mark task as completed

```sh 
tasklist-cli done 1
```

### Update task description

```sh
tasklist-cli update 1 "Buy organic groceries"
```

### Delete task

```sh
tasklist-cli delete 1
```

## Project Structure

```
.
├── cmd/          # Command implementations
├── methods/      # Core functionality 
├── models/       # Data structures
└── main.go       # Entry point
```

## Dependencies

- [github.com/spf13/cobra](https://github.com/spf13/cobra) v1.10.1 - CLI framework
- [github.com/spf13/pflag](https://github.com/spf13/pflag) v1.0.9
- [github.com/inconshreveable/mousetrap](https://github.com/inconshreveable/mousetrap) v1.1.0