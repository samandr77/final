package main

import (
	"final/internal/handlers"
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

	// Автоматическое создание таблиц (если еще не созданы)
	if err := db.AutoMigrate(&userService.User{}); err != nil {
		log.Fatal("Ошибка автоматической миграции:", err)
	}

	// Инициализация репозитория и сервиса для пользователей
	userRepo := userService.NewUserRepository(db)
	userSvc := userService.NewUserService(userRepo)
	userHandlers := handlers.NewUserHandlers(userSvc)

	// Настройка маршрутов
	e.GET("/users", userHandlers.GetUsers)
	e.POST("/users", userHandlers.PostUser)
	e.PATCH("/users/:id", userHandlers.PatchUserByID)
	e.DELETE("/users/:id", userHandlers.DeleteUserByID)

	// Запуск сервера
	e.Logger.Fatal(e.Start(":8080"))
}
