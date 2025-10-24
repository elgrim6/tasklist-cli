package main

import (
	"tasklist-cli/cmd"
)

func main() {
	const tasksFile = "tasks.json"
	cmd.Execute()
}
