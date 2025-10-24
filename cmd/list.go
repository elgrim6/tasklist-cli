package cmd

import (
	"fmt"
	"tasklist-cli/methods"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `Display all tasks in your task list.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load tasks from file
		tl, err := methods.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}
		if len(tl.Tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		// Print each task
		for _, task := range tl.Tasks {
			status := " "
			if task.Completed {
				status = "x"
			}
			fmt.Printf("[%s] %d: %s (Created: %s, Updated: %s)\n", status, task.ID, task.Description, task.CreatedAt.Format("2006-01-02 15:04:05"), task.UpdatedAt.Format("2006-01-02 15:04:05"))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Define flags for this command
}
