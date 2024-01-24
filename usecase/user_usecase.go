// usecase/user_usecase.go
package usecase

import (
	"github.com/fernandojr999/go-learning/domain"
	"github.com/fernandojr999/go-learning/repository"
	"golang.org/x/crypto/bcrypt"
)

// UserUsecase define as regras de negócios para usuários
type UserUsecase struct {
	userRepository repository.UserRepository
}

// NewUserUsecase cria uma instância de UserUsecase
func NewUserUsecase(userRepository repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

// CreateUser cria um novo usuário
func (u *UserUsecase) CreateUser(user *domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	return u.userRepository.CreateUser(user)
}

// AuthenticateUser verifica as credenciais do usuário
func (u *UserUsecase) AuthenticateUser(inputUser *domain.User) error {
	storedUser, err := u.userRepository.GetUserByUsername(inputUser.Username)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(inputUser.Password))
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) GetAllUsers() ([]domain.User, error) {
	return u.userRepository.GetAllUsers()
}
