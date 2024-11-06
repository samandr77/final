package taskService

import "gorm.io/gorm"

type TaskRepository struct {
	db *gorm.DB
}

func (r *TaskRepository) CreateTask(task *Task) (*Task, error) {
	if err := r.db.Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r *TaskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) GetTasksByUserID(userID uint, tasks *[]Task) error {
	if err := r.db.Where("user_id = ?", userID).Find(tasks).Error; err != nil {
		return err
	}
	return nil
}
