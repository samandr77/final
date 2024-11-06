package main

import (
	"final/internal/handlers"
	"final/internal/taskService"
	"final/internal/userService"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	e := echo.New() // Создаем новый экземпляр Echo

	// Настройка подключения к базе данных
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Ошибка подключения к базе данных")
	}

	// Автоматическое создание таблиц для пользователей и задач
	if err := db.AutoMigrate(&userService.User{}, &taskService.Task{}); err != nil {
		log.Fatal("Ошибка автоматической миграции:", err)
	}

	// Инициализация репозиториев и сервисов для пользователей и задач
	userRepo := userService.NewUserRepository(db)
	userSvc := userService.NewUserService(userRepo)
	userHandlers := handlers.NewUserHandlers(userSvc)

	taskRepo := taskService.NewTaskRepository(db)
	taskSvc := taskService.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskSvc)

	// Настройка маршрутов для пользователей
	e.GET("/users", userHandlers.GetUsers)
	e.POST("/users", userHandlers.PostUser)
	e.PATCH("/users/:id", userHandlers.PatchUserByID)
	e.DELETE("/users/:id", userHandlers.DeleteUserByID)

	// Настройка маршрутов для задач
	e.GET("/tasks", taskHandlers.GetTasksByUserID)
	e.POST("/tasks", taskHandlers.PostTask)

	// Запуск сервера
	e.Logger.Fatal(e.Start(":8080"))
}
