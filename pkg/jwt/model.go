package jwt

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	jwt.Payload
	Role string `json:"role"`
}
