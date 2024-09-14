package task

import (
	"errors"
	"fmt"
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
	return fmt.Errorf("invalid task status: %s", status)
}

func AddTask(description, filename string) error {
	taskFile, err := ReadTasks(filename)
	if err != nil {
		return fmt.Errorf("failed to add task: %w", err)
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

	if err := WriteTasks(filename, taskFile); err != nil {
		return fmt.Errorf("failed to save task: %w", err)
	}
	return nil
}

func UpdateTask(id int, description, filename string) error {
	taskFile, err := ReadTasks(filename)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	for i, task := range taskFile.Tasks {
		if task.Id == id {
			taskFile.Tasks[i].Description = description
			taskFile.Tasks[i].UpdatedAt = time.Now()
			if err := WriteTasks(filename, taskFile); err != nil {
				return fmt.Errorf("failed to save updated task: %w", err)
			}
			return nil
		}
	}

	return errors.New("task not found")
}

func DeleteTask(id int, filename string) error {
	taskFile, err := ReadTasks(filename)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	for i, task := range taskFile.Tasks {
		if task.Id == id {
			taskFile.Tasks = append(taskFile.Tasks[:i], taskFile.Tasks[i+1:]...)
			if err := WriteTasks(filename, taskFile); err != nil {
				return fmt.Errorf("failed to save tasks after deletion: %w", err)
			}
			return nil
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
		return fmt.Errorf("failed to mark task: %w", err)
	}

	for i, task := range taskFile.Tasks {
		if task.Id == id {
			taskFile.Tasks[i].Status = status
			taskFile.Tasks[i].UpdatedAt = time.Now()
			if err := WriteTasks(filename, taskFile); err != nil {
				return fmt.Errorf("failed to save marked task: %w", err)
			}
			return nil
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
		return nil, fmt.Errorf("failed to list tasks: %w", err)
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
