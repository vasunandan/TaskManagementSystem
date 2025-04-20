package dao

import (
    "gorm.io/gorm"
    "TaskManagementSystem/internal/models/postgres1"
	"TaskManagementSystem/pkg/logging"
	"fmt"
)

type TaskDAO struct {
    DB *gorm.DB
}


func (dao *TaskDAO) CreateTask(task *postgres1.Task) (*postgres1.Task, error) {
    result := dao.DB.Create(task)
    if result.Error != nil {
        return nil, result.Error
    }
    logging.Log(fmt.Sprintf("Created task with ID: %d", task.ID))
    return task, nil
}


func (dao *TaskDAO) GetAllTasks(page int, pageSize int, statusFilter string) ([]postgres1.Task, error) {
	offset := (page - 1) * pageSize
	var tasks []postgres1.Task
	query := dao.DB.Model(&postgres1.Task{})

	if statusFilter != "" {
		query = query.Where("status = ?", statusFilter)
	}

	result := query.Limit(pageSize).Offset(offset).Find(&tasks)
	return tasks, result.Error
}


func (dao *TaskDAO) GetTaskByID(id int) (*postgres1.Task, error) {
    var task postgres1.Task
    result := dao.DB.First(&task, id)
    if result.Error != nil {
        return nil, result.Error
    }
    logging.Log(fmt.Sprintf("Fetched task with ID: %d", id))
    return &task, nil
}


func (dao *TaskDAO) UpdateTask(id int, updatedTask *postgres1.Task) (*postgres1.Task, error) {
    var task postgres1.Task
    result := dao.DB.First(&task, id)
    if result.Error != nil {
        return nil, result.Error
    }

    task.Title = updatedTask.Title
    task.Description = updatedTask.Description
    task.Status = updatedTask.Status

    dao.DB.Save(&task)
    logging.Log(fmt.Sprintf("Updated task with ID: %d", id))
    return &task, nil
}


func (dao *TaskDAO) DeleteTask(id int) error {
    result := dao.DB.Delete(&postgres1.Task{}, id)
    if result.Error != nil {
        return result.Error
    }
    logging.Log(fmt.Sprintf("Deleted task with ID: %d", id))
    return nil
}