package cmd

import (
	"fmt"
	"tasklist-cli/methods"
	"tasklist-cli/models"
	"time"

	"github.com/spf13/cobra"
)

var desc string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `Add a new task to your task list.`,
	Run: func(cmd *cobra.Command, args []string) {
		desc = args[0]
		fmt.Printf("Task %s added to list\n", desc)

		tl, err := methods.LoadTasks()

		if err != nil {
			return
		}

		id := 1
		if len(tl.Tasks) > 0 {
			id = tl.Tasks[len(tl.Tasks)-1].ID + 1
		}

		task := models.Task{
			ID:          id,
			Description: desc,
			Completed:   false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		tl.Tasks = append(tl.Tasks, task)
		methods.SaveTasks(tl)
	},
}

func init() {
	// Add the 'hello' command to the root command
	rootCmd.AddCommand(addCmd)

	// Define a flag for this command
	addCmd.Flags().StringVarP(&desc, "add", "a", "New Task", "Description of the task to add")
}
