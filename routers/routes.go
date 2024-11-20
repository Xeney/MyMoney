package routers

import (
	"app/controllers"
	// session "app/middleware"

	// "github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// инициализирует маршруты, доступные до авторизации
func SetupBeforeUserRoutes(g *gin.RouterGroup) {
	routes := []struct {
		method      string
		path        string
		handlerFunc gin.HandlerFunc
	}{
		{"GET", "/", controllers.HomeGetHanler},
		{"GET", "/secure", controllers.SecureGetHandler},
		{"GET", "/feedback", controllers.FeedbackGetHandler},
		{"GET", "/help", controllers.HelpGetHandler},
		{"GET", "/error", controllers.ErrorGetHandler},
		{"GET", "/sign-in", controllers.LoginGetHandler},
		{"GET", "/sign-up", controllers.RegistrationGetHandler},
	}

	for _, route := range routes {
		switch route.method {
		case "GET":
			g.GET(route.path, route.handlerFunc)
		}
	}
}

// SetupAfterUserRoutes инициализирует маршруты, доступные после авторизации
func SetupAfterUserRoutes(g *gin.RouterGroup) {
	routes := []struct {
		method      string
		path        string
		handlerFunc gin.HandlerFunc
	}{
		// {"GET", "/dashboard", controllers.DashboardGetHandler},
		// {"POST", "/logout", controllers.LogoutPostHandler},
	}

	for _, route := range routes {
		switch route.method {
		case "GET":
			g.GET(route.path, route.handlerFunc)
		case "POST":
			g.POST(route.path, route.handlerFunc)
			// Добавление других методов при необходимости
		}
	}
}
