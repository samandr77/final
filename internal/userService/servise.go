package userService

import (
	"errors"
	"gorm.io/gorm"
)

// UserService - структура, отвечающая за логику работы с пользователями.
type UserService struct {
	repo *UserRepository
}

// GetUserTasks - метод для получения всех задач пользователя.
func GetUserTasks(db *gorm.DB, userID uint) ([]Task, error) { // Используем Task из userService
	var tasks []Task
	err := db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]User, error) { // Используем User из userService
	return s.repo.GetAllUsers()
}

func (s *UserService) CreateUser(user *User) (*User, error) { // Используем User из userService
	if user.Name == "" || user.Email == "" {
		return nil, errors.New("поля 'name' и 'email' обязательны")
	}
	return s.repo.CreateUser(user)
}

// Обновление пользователя по ID
func (s *UserService) UpdateUserByID(id string, updatedUser *User) (*User, error) { // Используем User из userService
	if updatedUser.Name == "" || updatedUser.Email == "" {
		return nil, errors.New("поля 'name' и 'email' обязательны")
	}
	return s.repo.UpdateUserByID(id, updatedUser)
}

// Удаление пользователя по ID
func (s *UserService) DeleteUserByID(id int) error {
	return s.repo.DeleteUserByID(id)
}

// Получение задач для конкретного пользователя
func (s *UserService) GetTasksForUser(userID uint) ([]Task, error) {
	var tasks []Task
	// Используем репозиторий для получения всех задач пользователя
	if err := s.repository.GetTasksByUserID(userID, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
