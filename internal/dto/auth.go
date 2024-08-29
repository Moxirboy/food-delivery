package dto

type Tokens struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Token     *Tokens `json:"tokens"`
}

type SignUpRequest struct {
	FisrtName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
type CheckRequest struct {
	AccessToken string `json:"access_token"`
}

type CheckResponse struct {
	ID   string `json:"id"`
	Role string `json:"role"`
}

type RenewRequest struct {
	RefreshToken string `json:"refresh_token"`
}
