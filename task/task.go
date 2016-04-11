package task

import (
	"errors"
)

var NoTaskError = errors.New("No task with such id.")

type Task struct {
	Id       int    `json:id`
	Text     string `json:text`
	Complete bool   `json:complete`
}

type Storage struct {
	tasks  []Task `json:tasks`
	lastId int    `json:lastId`
}

func NewStorage() Storage {
	tasks := make([]Task, 0)

	return Storage{
		tasks:  tasks,
		lastId: 0,
	}
}

func newTask(id int, desc string, complete bool) Task {
	return Task{
		Id:       id,
		Text:     desc,
		Complete: complete,
	}
}

func (s *Storage) incrementId() {
	s.lastId++
}

func (s *Storage) AddTask(desc string) int {
	defer s.incrementId()
	s.tasks = append(s.tasks, newTask(s.lastId, desc, false))

	return s.lastId
}

func (s Storage) GetTasks() []Task {
	return s.tasks
}

func (s Storage) GetTask(id int) (Task, error) {
	for _, task := range s.tasks {
		if task.Id == id {
			return task, nil
		}
	}

	return Task{}, NoTaskError
}

func (s *Storage) UpdateTask(id int, desc string, complete bool) (Task, error) {
	for i, task := range s.tasks {
		if s.tasks[i].Id == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			s.tasks = append(s.tasks, newTask(id, desc, complete))
			return task, nil
		}
	}

	return Task{}, NoTaskError
}

func (s *Storage) DeleteTask(id int) (Task, error) {
	for i, task := range s.tasks {
		if task.Id == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return task, nil
		}
	}

	return Task{}, NoTaskError
}
