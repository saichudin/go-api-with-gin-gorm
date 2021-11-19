package user

type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
