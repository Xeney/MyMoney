package globals

// Структура пользователя
type Feedback struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	MessageTheme string `json:"message_theme"`
	Message      string `json:"message"`
}
