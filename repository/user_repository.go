// repository/user_repository.go
package repository

import (
	"database/sql"

	"github.com/fernandojr999/go-learning/domain"
)

// UserRepository define as operações do banco de dados para usuários
type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByUsername(username string) (*domain.User, error)
	GetAllUsers() ([]domain.User, error)
}

// userRepo implementa a interface UserRepository
type userRepo struct {
	db *sql.DB
}

// NewUserRepository cria uma instância de userRepo
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{
		db: db,
	}
}

// CreateUser insere um novo usuário no banco de dados
func (r *userRepo) CreateUser(user *domain.User) error {
	_, err := r.db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
	return err
}

// GetUserByUsername busca um usuário pelo nome de usuário
func (r *userRepo) GetUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := r.db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password)
	return &user, err
}

func (r *userRepo) GetAllUsers() ([]domain.User, error) {
	rows, err := r.db.Query("SELECT id, username, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
