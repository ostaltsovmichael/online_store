package user_services

type User struct {
	Id       uint64 `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
type Err_mes struct {
	Err_mes string `json:"err_mes"`
}
