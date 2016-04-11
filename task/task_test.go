package task

import (
	"testing"
)

func TestNewStorage(t *testing.T) {
	s := NewStorage()

	if len(s.tasks) != 0 {
		t.Error("non empty storage initialized.")
	}

	if s.lastId != 0 {
		t.Error("lastId initialized with offset.")
	}
}

func TestAddNewTask(t *testing.T) {
	taskText := "new test task"

	s := NewStorage()
	id := s.AddTask(taskText)

	if len(s.tasks) != 1 {
		t.Errorf("error on task appending. Tasks storage: %s.", s.tasks)
	}

	if s.lastId != 1 {
		t.Error("lastId didn't increment during new task adding.")
	}

	if s.tasks[0].Text != taskText || s.tasks[0].Id != id || s.tasks[0].Complete {
		t.Errorf("Wrong task created. Current task: %s.", s.tasks[0])
	}
}

func TestFindTask(t *testing.T) {
	taskText := "new test task"

	s := NewStorage()
	id := s.AddTask(taskText)

	task, err := s.GetTask(id)
	if err != nil {
		t.Error(err)
	}

	if task.Text != taskText {
		t.Errorf("wrong task found: %s", task)
	}
}

func TestUpdateTask(t *testing.T) {
	taskText := "new test task"
	newTaskText := "changed description"

	s := NewStorage()
	id := s.AddTask(taskText)

	task, err := s.UpdateTask(id, newTaskText, true)
	if err != nil {
		t.Error(err)
	}

	if task.Text != newTaskText || !task.Complete {
		t.Errorf("wrong task change: %s", task)
	}
}

func TestDeleteTask(t *testing.T) {
	taskText := "new test task"

	s := NewStorage()
	id := s.AddTask(taskText)

	_, err := s.DeleteTask(id)
	if err != nil {
		t.Error(err)
	}

	if len(s.tasks) != 0 {
		t.Error("Task wasn't deleted.")
	}
}
