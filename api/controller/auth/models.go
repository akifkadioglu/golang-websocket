package auth

type RegisterBody struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
