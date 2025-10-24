package methods

import (
	"encoding/json"
	"os"
	"tasklist-cli/models"
)

const tasksFile = "tasks.json"

func LoadTasks() (models.TaskList, error) {
	var tl models.TaskList

	data, err := os.ReadFile(tasksFile)
	if err != nil {
		if os.IsNotExist(err) {
			return tl, nil
		}
		return tl, err
	}

	err = json.Unmarshal(data, &tl)
	return tl, err
}

func SaveTasks(tl models.TaskList) error {
	data, err := json.MarshalIndent(tl, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(tasksFile, data, 0644)
}
