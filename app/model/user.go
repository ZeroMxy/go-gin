package model

type User struct {
	Model
	UserType int    `json:"userType"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Gender   int    `json:"gender"`
	Age      int    `json:"age"`
	Avatar   string `json:"avatar"`
	Status   int    `json:"status"`
}
