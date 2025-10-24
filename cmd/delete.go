package cmd

import (
	"fmt"
	"log"
	"strconv"
	"tasklist-cli/methods"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [task ID]",
	Short: "Delete a task",
	Long:  `Delete a specific task from your task list using its ID.`,
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
				tl.Tasks[i] = tl.Tasks[len(tl.Tasks)-1]
				tl.Tasks = tl.Tasks[:len(tl.Tasks)-1]
				found = true
				fmt.Printf("Task %d deleted\n", id)
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
	rootCmd.AddCommand(deleteCmd)
}
