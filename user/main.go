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
	Password  string
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
}

func NewUserService(repo Repository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *User) error {
	// validate user data

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)
	user.ID = uuid.NewString()
	now := time.Now().Format("2006-01-02 15:04:05")
	user.CreatedAt = now
	user.UpdatedAt = now

	return s.repo.Save(user)
}

func (s *UserService) GetUserById(id string) (*User, error) {
	return s.repo.GetById(id)
}

func (s *UserService) GetUserByEmail(email string) (*User, error) {
	return s.repo.GetByEmail(email)
}

func (s *UserService) Update(user User) error {
	return s.repo.Update(&user)
}

func (s *UserService) GetAllUsers() ([]*User, error) {
	return s.repo.List()
}