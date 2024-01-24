// domain/user.go
package domain

// User struct representa um usuário
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
