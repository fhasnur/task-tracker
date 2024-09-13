package task

import (
	"errors"
	"time"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

const (
	StatusTodo       = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
)

func ValidateStatus(status string) error {
	validStatuses := []string{"all", StatusTodo, StatusInProgress, StatusDone}
	for _, validStatus := range validStatuses {
		if status == validStatus {
			return nil
		}
	}
	return errors.New("invalid task status")
}

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

func DeleteTask(id int, filename string) error {
	taskFile, err := ReadTasks(filename)
	if err != nil {
		return err
	}

	for i, task := range taskFile.Tasks {
		if task.Id == id {
			taskFile.Tasks = append(taskFile.Tasks[:i], taskFile.Tasks[i+1:]...)
			return WriteTasks(filename, taskFile)
		}
	}
	return errors.New("task not found")
}

func MarkTask(id int, status, filename string) error {
	if err := ValidateStatus(status); err != nil {
		return err
	}

	taskFile, err := ReadTasks(filename)
	if err != nil {
		return err
	}

	for i, task := range taskFile.Tasks {
		if task.Id == id {
			taskFile.Tasks[i].Status = status
			taskFile.Tasks[i].UpdatedAt = time.Now()
			return WriteTasks(filename, taskFile)
		}
	}
	return errors.New("task not found")
}

func ListTasks(status, filename string) ([]Task, error) {
	if err := ValidateStatus(status); err != nil {
		return nil, err
	}

	taskFile, err := ReadTasks(filename)
	if err != nil {
		return nil, err
	}

	var filteredTasks []Task
	switch status {
	case "all":
		filteredTasks = taskFile.Tasks
	case StatusTodo:
		for _, task := range taskFile.Tasks {
			if task.Status == StatusTodo {
				filteredTasks = append(filteredTasks, task)
			}
		}
	case StatusInProgress:
		for _, task := range taskFile.Tasks {
			if task.Status == StatusInProgress {
				filteredTasks = append(filteredTasks, task)
			}
		}
	case StatusDone:
		for _, task := range taskFile.Tasks {
			if task.Status == StatusDone {
				filteredTasks = append(filteredTasks, task)
			}
		}
	}

	if len(filteredTasks) == 0 {
		return nil, errors.New("no tasks found")
	}

	return filteredTasks, nil
}
