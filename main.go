package main

import (
	// "net/http"
	"app/middlewares"
	"app/routers"
	"log"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация Gin
	router := gin.Default()

	// Настройка сессий
	store := middlewares.SetSession()

	// Группы маршрутов
	beforeAuthorization := router.Group("/")
	afterAuthorization := router.Group("/")

	// Применение middleware для групп маршрутов
	beforeAuthorization.Use(sessions.Sessions("session-name", store))
	afterAuthorization.Use(sessions.Sessions("session-name", store))
	afterAuthorization.Use(middlewares.AuthSession)

	// Регистрация маршрутов
	routers.SetupBeforeUserRoutes(beforeAuthorization)
	routers.SetupAfterUserRoutes(afterAuthorization)

	// Load files
	router.LoadHTMLGlob("./ui/templates/*")

	// Настройка статических файлов
	staticFiles := map[string]string{
		"style.css":           "./ui/static/css/style.css",
		"background-home.jpg": "./ui/static/images/background-home.jpg",
		"back-money.jpg":      "./ui/static/images/home/back-money.jpg",
		"money-logo.jpg":      "./ui/static/images/home/money-logo.jpg",
		"icon-accent.svg":     "./ui/static/images/icons/icon-accent.svg",
		"message.svg":         "./ui/static/images/icons/message.svg",
		"phone.svg":           "./ui/static/images/icons/phone.svg",
		"pos.svg":             "./ui/static/images/icons/pos.svg",
		"time.svg":            "./ui/static/images/icons/time.svg",
	}

	// Регистрация статических файлов
	for route, file := range staticFiles {
		router.StaticFile(route, file)
	}

	// Запуск сервера
	if err := router.Run(":5000"); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
