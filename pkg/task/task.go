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
	switch status {
	case StatusTodo, StatusInProgress, StatusDone:
		return nil
	default:
		return errors.New("invalid task status")
	}
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

func ListAllTasks(filename string) ([]Task, error) {
	taskFile, err := ReadTasks(filename)
	if err != nil {
		return nil, err
	}

	var taskList []Task
	for _, task := range taskFile.Tasks {
		taskList = append(taskList, task)
	}

	return taskList, nil
}
