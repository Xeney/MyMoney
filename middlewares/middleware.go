package middlewares

import (
	"app/globals"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// SetSession устанавливает новое хранилище для сохранения сессий
func SetSession() sessions.Store {
	store := sessions.NewCookieStore([]byte("my-secret-key"))
	return store
}

// AuthSession - промежуточное ПО для авторизации пользователей через сессии
func AuthSession(c *gin.Context) {
	session := sessions.Default(c)
	sessionID := session.Get(globals.SessionKey)
	if sessionID == nil {
		c.Redirect(http.StatusFound, "/error") // Перенаправляем пользователя на страницу ошибки
		c.Abort()                              // Прерываем цепочку обработчиков
		return
	}
	c.Next() // Продолжаем выполнение следующего обработчика
}

// SaveSession сохраняет сессию пользователя
func SaveSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Set(globals.SessionKey, true) // Устанавливаем значение сессии
	if err := session.Save(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err) // Обработка ошибок при сохранении
		return
	}
}

// ClearSession очищает сессию пользователя
func ClearSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(globals.SessionKey) // Удаляем элемент сессии
	if err := session.Save(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err) // Обработка ошибок при сохранении
		return
	}
}

// GetSession возвращает значение сессии для пользователя
func GetSession(c *gin.Context) interface{} {
	session := sessions.Default(c)
	return session.Get(globals.SessionKey) // Возвращаем значение сессии
}

// CheckSession проверяет, существует ли сессия пользователя
func CheckSession(c *gin.Context) bool {
	session := sessions.Default(c)
	return session.Get(globals.SessionKey) != nil // Возвращаем истинное, если сессия существует
}
