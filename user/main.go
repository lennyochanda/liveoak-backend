package user

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

)

type User struct {
	ID        string
	Username  string
	Email     string
	PasswordHash  string
	CreatedAt string
	UpdatedAt string
}

type Repository interface {
	Save(user *User) error
	GetById(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(user *User) error
	List() ([]*User, error)
}

type UserService struct {
	repo Repository
	logger Logger
}


func NewUserService(repo Repository) *UserService {
	userService := UserService {
		repo: repo,
		logger: Logger{"user-service.log.txt"},
	}

	userService.logger.Log("Created New User Service", nil)
	return &userService
}

func (s *UserService) CreateUser(userName, email, password string) error {
	var user User

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}


	user.PasswordHash = string(bytes)
	user.ID = uuid.NewString()
	now := time.Now().Format("2006-01-02 15:04:05")
	user.CreatedAt = now
	user.UpdatedAt = now

	s.logger.Log("Creating New User: ", user)
	return s.repo.Save(&user)
}

func (s *UserService) GetUserById(id string) (*User, error) {
	s.logger.Log("Getting User with ID: ", id)
	return s.repo.GetById(id)
}

func (s *UserService) GetUserByEmail(email string) (*User, error) {
	s.logger.Log("Getting User with Email: ", email)
	return s.repo.GetByEmail(email)
}

func (s *UserService) Update(user User) error {
	s.logger.Log("Updating User with Details: ", user)
	return s.repo.Update(&user)
}

func (s *UserService) GetAllUsers() ([]*User, error) {
	s.logger.Log("Getting List of Users", nil)
	return s.repo.List()
}

