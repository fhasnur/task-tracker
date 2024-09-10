package task

import (
	"encoding/json"
	"errors"
	"os"
)

type TaskFile struct {
	LastId int    `json:"last_id"`
	Tasks  []Task `json:"tasks"`
}

func ReadTasks(filename string) (TaskFile, error) {
	var taskFile TaskFile

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return TaskFile{LastId: 0, Tasks: []Task{}}, nil
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		return taskFile, err
	}

	err = json.Unmarshal(file, &taskFile)
	if err != nil {
		return taskFile, err
	}

	return taskFile, nil
}

func WriteTasks(filename string, taskFile TaskFile) error {
	fileData, err := json.MarshalIndent(taskFile, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, fileData, 0644)
	if err != nil {
		return err
	}

	return nil
}
