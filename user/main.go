package user

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/lennyochanda/LiveOak/logger"
	"github.com/lennyochanda/LiveOak/types"
)

type Repository interface {
	Save(user *types.User) error
	GetById(id string) (*types.User, error)
	GetByEmail(email string) (*types.User, error)
	Update(user *types.User) error
	List() ([]*types.User, error)
}

type UserService struct {
	repo   Repository
	logger logger.Logger
}

func NewUserService(repo Repository) *UserService {
	userService := UserService{
		repo:   repo,
		logger: logger.Logger{FileName: "user-service.log.txt"},
	}

	userService.logger.Log("Created New User Service", nil)
	return &userService
}

func (s *UserService) CreateUser(userdto CreateUserDTO) error {
	var user types.User

	bytes, err := bcrypt.GenerateFromPassword([]byte(userdto.Password), 14)
	if err != nil {
		return err
	}

	user.Username = userdto.UserName
	user.Email = userdto.Email
	user.PasswordHash = string(bytes)
	user.ID = uuid.NewString()
	now := time.Now().Format("2006-01-02 15:04:05")
	user.CreatedAt = now
	user.UpdatedAt = now

	s.logger.Log("Creating New User: ", user)
	return s.repo.Save(&user)
}

func (s *UserService) GetUserById(id string) (*types.User, error) {
	s.logger.Log("Getting User with ID: ", id)
	return s.repo.GetById(id)
}

func (s *UserService) GetUserByEmail(email string) (*types.User, error) {
	s.logger.Log("Getting User with Email: ", email)
	return s.repo.GetByEmail(email)
}

func (s *UserService) Update(user types.User) error {
	s.logger.Log("Updating User with Details: ", user)
	return s.repo.Update(&user)
}

func (s *UserService) GetAllUsers() ([]*types.User, error) {
	s.logger.Log("Getting List of Users", nil)
	return s.repo.List()
}

func (s *UserService) CheckPassword(loginForm LoginUserDTO) (bool, error) {
	user, err := s.repo.GetByEmail(loginForm.Email)
	if err != nil {
		return false, err
	}
	result := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginForm.Password))

	if result == nil {
		return true, nil
	} else {
		return false, nil
	}
}
