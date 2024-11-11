package routers

import (
	"app/controllers"
	// session "app/middleware"

	// "github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func BeforeUserRouter(g *gin.RouterGroup) {
	g.GET("/", controllers.HomeGetHanler)
	g.GET("/secure", controllers.SecureGetHandler)
	g.GET("/feedback", controllers.FeedbackGetHandler)
	g.GET("/help", controllers.HelpGetHandler)
	g.GET("/error", controllers.ErrorGetHandler)
	g.GET("/sign-in", controllers.LoginGetHandler)
	g.GET("/sign-up", controllers.RegistrationGetHandler)
}

func AfterUserRouter(g *gin.RouterGroup) {
	// g.GET("/dashboard", controllers.DashboardGetHandler)
	// g.GET("/logout", controllers.LogoutPostHandler)
}
