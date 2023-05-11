package user

import (
	"database/sql"
	"fmt"
	"log"
)

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