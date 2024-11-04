package handlers

import (
	"final/internal/userService"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserHandlers struct {
	UserService *userService.UserService
}

// Конструктор для UserHandlers
func NewUserHandlers(userService *userService.UserService) *UserHandlers {
	return &UserHandlers{UserService: userService}
}

// Получение всех пользователей
func (h *UserHandlers) GetUsers(ctx echo.Context) error {
	users, err := h.UserService.GetAllUsers()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "Ошибка при получении пользователей"})
	}
	return ctx.JSON(http.StatusOK, users)
}

// Создание пользователя
func (h *UserHandlers) PostUser(ctx echo.Context) error {
	var user userService.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Ошибка при разборе данных пользователя"})
	}
	createdUser, err := h.UserService.CreateUser(&user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "Ошибка при создании пользователя"})
	}
	return ctx.JSON(http.StatusCreated, createdUser)
}

// Обновление пользователя по ID
func (h *UserHandlers) PatchUserByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Некорректный ID пользователя"})
	}

	var userUpdates userService.User
	if err := ctx.Bind(&userUpdates); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Ошибка при разборе данных пользователя"})
	}

	updatedUser, err := h.UserService.UpdateUserByID(strconv.Itoa(id), &userUpdates)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "Ошибка при обновлении пользователя"})
	}

	return ctx.JSON(http.StatusOK, updatedUser)
}

// Удаление пользователя по ID
func (h *UserHandlers) DeleteUserByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Некорректный ID пользователя"})
	}

	if err := h.UserService.DeleteUserByID(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "Ошибка при удалении пользователя"})
	}

	return ctx.NoContent(http.StatusNoContent)
}
