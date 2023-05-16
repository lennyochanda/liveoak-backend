package user

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/lennyochanda/LiveOak/logger"
)

type MySQLUserRepository struct {
	db *sql.DB
	logger logger.Logger
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db: db, logger: logger.Logger{FileName: "user-repository.log.txt"},}
}

func (r *MySQLUserRepository) Save(user *User) error {
	stmt, err := r.db.Prepare(`INSERT INTO users(id, username, email, password, createdAt, updatedAt) VALUES(?, ?, ?, ?, ?, ?)`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	if _, err := stmt.Exec(user.ID, user.Username, user.Email, user.PasswordHash, user.CreatedAt, user.UpdatedAt); err != nil {
		return fmt.Errorf("could not execute query: %w", err)
	}

	r.logger.Log("Insert into Users: ", user)
	return nil
}

func (r *MySQLUserRepository) GetById(id string) (*User, error) {
	user := &User{}
	err := r.db.QueryRow("SELECT id, username, email, password, createdAt, updatedAt FROM users where id=?", id).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	r.logger.Log("Get User by Id: ", id)
	return user, nil
}

func (r *MySQLUserRepository) GetByEmail(email string) (*User, error) {
	user := &User{}
	err := r.db.QueryRow("SELECT id, username, email, password, createdAt, updatedAt FROM users where email=?", email).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	r.logger.Log("Get User by Email: ", email)
	return user, nil
}

func (r *MySQLUserRepository) Update(user *User) error {
	stmt, err := r.db.Prepare("UPDATE users SET name=?, email=?, password=? WHERE id=?")
	if err != nil {
		return err
	}

	defer stmt.Close()
	
	_, err = stmt.Exec(user.Username, user.Email, user.PasswordHash, user.UpdatedAt, user.ID)
	
	if err != nil {
		return err
	}

	r.logger.Log("Update User: ", user)
	return nil
}

func (r *MySQLUserRepository) List() ([]*User, error) {
	var users []*User

	rows, err := r.db.Query("SELECT id, username, email, password, createdAt, updatedAt FROM users")
	if err != nil {
		return nil, err
	}

	
	defer rows.Close()
	
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	r.logger.Log("\n==========User List==========\n", nil)
	for index := range users {
		r.logger.Log(strconv.Itoa(index), *users[index])
	}
	return users, nil
}