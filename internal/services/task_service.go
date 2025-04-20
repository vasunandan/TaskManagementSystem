package services

import (
	"TaskManagementSystem/internal/domains/dao"
	"TaskManagementSystem/internal/models/postgres1"
)

type TaskService interface {
	GetAllTasks(page int, pageSize int, statusFilter string) ([]postgres1.Task, error)
	GetTaskByID(id int) (*postgres1.Task, error)
	CreateTask(task *postgres1.Task) (*postgres1.Task, error)
	UpdateTask(id int,task *postgres1.Task) (*postgres1.Task, error)
	DeleteTask(id int) error
}

type TaskServiceImpl struct {
	TaskDAO *dao.TaskDAO
}

func (s *TaskServiceImpl) GetAllTasks(page int, pageSize int, statusFilter string) ([]postgres1.Task, error) {
	return s.TaskDAO.GetAllTasks(page, pageSize , statusFilter)
}

func (s *TaskServiceImpl) GetTaskByID(id int) (*postgres1.Task, error) {
	return s.TaskDAO.GetTaskByID(id)
}

func (s *TaskServiceImpl) CreateTask(task *postgres1.Task) (*postgres1.Task, error) {
	return s.TaskDAO.CreateTask(task)
}

func (s *TaskServiceImpl) UpdateTask(id int,task *postgres1.Task) (*postgres1.Task, error) {
	return s.TaskDAO.UpdateTask(id,task)
}

func (s *TaskServiceImpl) DeleteTask(id int) error {
	return s.TaskDAO.DeleteTask(id)
}