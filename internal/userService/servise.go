package userService

import (
	"errors"
)

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) CreateUser(user *User) (*User, error) {
	if user.Name == "" || user.Email == "" {
		return nil, errors.New("поля 'name' и 'email' обязательны")
	}
	return s.repo.CreateUser(user)
}

// Обновление пользователя по ID
func (s *UserService) UpdateUserByID(id string, updatedUser *User) (*User, error) {
	if updatedUser.Name == "" || updatedUser.Email == "" {
		return nil, errors.New("поля 'name' и 'email' обязательны")
	}
	return s.repo.UpdateUserByID(id, updatedUser)
}

// Удаление пользователя по ID
func (s *UserService) DeleteUserByID(id int) error {
	return s.repo.DeleteUserByID(id)
}
