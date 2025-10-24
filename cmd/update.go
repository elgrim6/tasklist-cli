package cmd

import (
	"fmt"
	"log"
	"strconv"
	"tasklist-cli/methods"
	"time"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [task ID] [new description]",
	Short: "Update description of a task",
	Long:  `Update the description of a specific task in your task list using its ID.`,
	Args:  cobra.ExactArgs(2),
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
				tl.Tasks[i].Description = args[1]
				tl.Tasks[i].UpdatedAt = time.Now()
				found = true
				fmt.Printf("Task %d updated\n", id)
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
	rootCmd.AddCommand(updateCmd)
}
