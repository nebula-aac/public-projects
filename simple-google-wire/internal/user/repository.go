package user

import (
	"context"
	"database/sql"

	"github.com/nebula-aac/public-projects/simple-google-wire/internal/domain"
)

type repository struct {
	db *sql.DB
}

func (r *repository) FindByUsername(ctx context.Context, username string) (*domain.UserEntity, error) {
	var user domain.UserEntity
	err := r.db.QueryRowContext(ctx, "SELECT id, username, password, email FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
