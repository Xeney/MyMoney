package main

import (
	// "net/http"
	"app/middlewares"
	"app/routers"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	store := middlewares.SetSession()

	// Initializing group router
	beforeAuthorization := router.Group("/")
	afterAuthorization := router.Group("/")

	// Using middleware
	beforeAuthorization.Use(sessions.Sessions("session-name", store))
	afterAuthorization.Use(sessions.Sessions("session-name", store))
	afterAuthorization.Use(middlewares.AuthSession)

	// Routing
	routers.BeforeUserRouter(beforeAuthorization)
	routers.AfterUserRouter(afterAuthorization)

	// Load files
	router.LoadHTMLGlob("./ui/templates/*")
	router.StaticFile("style.css", "./ui/static/css/style.css")

	router.StaticFile("background-home.jpg", "./ui/static/images/background-home.jpg")
	router.StaticFile("back-money.jpg", "./ui/static/images/home/back-money.jpg")
	router.StaticFile("money-logo.jpg", "./ui/static/images/home/money-logo.jpg")
	router.StaticFile("icon-accent.svg", "./ui/static/images/icons/icon-accent.svg")
	router.StaticFile("message.svg", "./ui/static/images/icons/message.svg")
	router.StaticFile("phone.svg", "./ui/static/images/icons/phone.svg")
	router.StaticFile("pos.svg", "./ui/static/images/icons/pos.svg")
	router.StaticFile("time.svg", "./ui/static/images/icons/time.svg")

	// Start server
	router.Run(":5000")
}
