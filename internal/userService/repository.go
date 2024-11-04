package userService

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) GetAllUsers() ([]User, error) {
	var users []User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) CreateUser(user *User) (*User, error) {
	if err := repo.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Обновление пользователя по ID
func (repo *UserRepository) UpdateUserByID(id string, updatedUser *User) (*User, error) {
	var user User
	if err := repo.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	// Обновляем поля пользователя
	user.Name = updatedUser.Name
	user.Email = updatedUser.Email

	if err := repo.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Удаление пользователя по ID
func (repo *UserRepository) DeleteUserByID(id int) error {
	if err := repo.db.Delete(&User{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
