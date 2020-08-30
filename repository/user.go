package repository

import (
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hatena/go-Intern-Diary/model"
)

var userNotFoundError = model.NotFoundError("user")

func (r *repository) CreateNewUser(name string, passwordHash string) error {
	id, err := r.generateID()
	if err != nil {
		return err
	}
	now := time.Now()
	_, err = r.db.Exec(
		`INSERT INTO user
			(id, name, password_hash, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?)`,
		id, name, passwordHash, now, now,
	)
	return err
}

func (r *repository) FindUserByName(name string) (*model.User, error) {
	var user model.User
	err := r.db.Get(
		&user,
		`SELECT id,name FROM user
			WHERE name = ? LIMIT 1`, name,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, userNotFoundError
		}
		return nil, err
	}
	return &user, nil
}

func (r *repository) CreateNewToken(userID uint64, token string, expiresAt time.Time) error {
	now := time.Now()
	_, err := r.db.Exec(
		`INSERT INTO user_session
			(user_id, token, expires_at, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?)`,
		userID, token, expiresAt, now, now,
	)
	return err
}
