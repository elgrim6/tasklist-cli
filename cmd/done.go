package cmd

import (
	"fmt"
	"log"
	"strconv"
	"tasklist-cli/methods"
	"time"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [task ID]",
	Short: "Mark a task as done",
	Long:  `Mark a specific task in your task list as completed using its ID.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tl, err := methods.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("failed to convert string to integer: %v", err)
		}
		found := false

		for i := range tl.Tasks {
			if tl.Tasks[i].ID == id {
				tl.Tasks[i].Completed = true
				tl.Tasks[i].UpdatedAt = time.Now()
				found = true
				fmt.Printf("Task %d marked as done\n", id)
				break
			}
		}

		if !found {
			fmt.Printf("Task %d not found\n", id)
		}
		methods.SaveTasks(tl)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Define flags for this command
	//doneCmd.Flags().StringVarP(&desc, "add", "a", "New Task", "Description of the task to add")
}
