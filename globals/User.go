package globals

// Структура пользователя
type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
