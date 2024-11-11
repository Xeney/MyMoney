package middlewares

import (
	"app/globals"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Set new store for save session
func SetSession() sessions.Store {
	store := sessions.NewCookieStore([]byte("my-secret-key"))
	return store
}

// User Auth Session Middle
func AuthSession(c *gin.Context) {
	session := sessions.Default(c)
	sessionID := session.Get(globals.SessionKey)
	if sessionID == nil {
		c.Redirect(http.StatusFound, "/error")
		return
	}
	c.Next()
}

// Save Session for User
func SaveSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Set(globals.SessionKey, true)
	session.Save()
}

// Clear Session for User
func ClearSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(globals.SessionKey)
	err := session.Save()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

// Get Session for User
func GetSession(c *gin.Context) interface{} {
	session := sessions.Default(c)
	sessionID := session.Get(globals.SessionKey)
	return sessionID
}

// Check Session for User
func CheckSession(c *gin.Context) bool {
	session := sessions.Default(c)
	sessionID := session.Get(globals.SessionKey)
	return sessionID != nil
}
