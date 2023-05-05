package user

import (
	"database/sql"
	"fmt"
	"log"
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

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) Save(user *User) error {
	// validate user data
	// save user to MySQL db
	stmt, err := r.db.Prepare(`INSERT INTO users(id, username, email, password, createdAt, updatedAt) VALUES(?, ?, ?, ?, ?, ?)`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	if _, err := stmt.Exec(user.ID, user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt); err != nil {
		return fmt.Errorf("could not execute query: %w", err)
	}

	return nil
}

func (r *MySQLUserRepository) GetById(id string) (*User, error) {
	// get user from MySQL by id

	row := r.db.QueryRow("SELECT id, username, email, password, createdAt, updatedAt FROM users where id=?", id)
	user := &User{}

	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *MySQLUserRepository) GetByEmail(email string) (*User, error) {
	row := r.db.QueryRow("SELECT id, username, email, password, createdAt, updatedAt FROM users where email=?", email)
	user := &User{}

	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *MySQLUserRepository) Update(user *User) error {
	query := `UPDATE users SET name=?, email=?, password=? WHERE id=?`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Email, user.Password, user.UpdatedAt, user.ID)

	if err != nil {
		return err
	}
	return nil
}

func (r *MySQLUserRepository) List() ([]*User, error) {
	var users []*User

	query := `SELECT id, username, email, password, createdAt, updatedAt FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}