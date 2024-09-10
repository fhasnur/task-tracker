package task

import (
	"errors"
	"time"
)

func AddTask(description, filename string) error {
	taskFile, err := ReadTasks(filename)
	if err != nil {
		return err
	}

	newTask := Task{
		Id:          taskFile.LastId + 1,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	taskFile.LastId++
	taskFile.Tasks = append(taskFile.Tasks, newTask)

	return WriteTasks(filename, taskFile)
}

func UpdateTask(id int, description, filename string) error {
	taskFile, err := ReadTasks(filename)
	if err != nil {
		return err
	}

	for i, task := range taskFile.Tasks {
		if task.Id == id {
			taskFile.Tasks[i].Description = description
			taskFile.Tasks[i].UpdatedAt = time.Now()
			return WriteTasks(filename, taskFile)
		}
	}

	return errors.New("task not found")
}
