package domain

type AuthRequest struct {
	Token string `json:"token"`
}

type GoogleUser struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Sub   string `json:"sub"`
}
