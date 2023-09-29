package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Ahmed-Sermani/driver-tracking-system/users/internal/domain"
)

type UserRepository struct {
	tableName string
	db        *sql.DB
}

var _ domain.UserRepository = (*UserRepository)(nil)

func NewUserRepository(tableName string, db *sql.DB) UserRepository {
	return UserRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r UserRepository) Save(ctx context.Context, user *domain.User) error {
	const query = "INSERT INTO %s (id, name, phone, enabled) VALUES ($1, $2, $3, $4)"

	_, err := r.db.ExecContext(ctx, r.table(query), user.ID, user.Name, user.Phone, user.Enabled)

	return err
}

func (r UserRepository) Find(ctx context.Context, userID string) (*domain.User, error) {
	const query = "SELECT name, phone, enabled FROM %s WHERE id = $1 LIMIT 1"

	user := &domain.User{
		ID: userID,
	}

	err := r.db.QueryRowContext(ctx, r.table(query), userID).Scan(&user.Name, &user.Phone, &user.Enabled)

	return user, err
}

func (r UserRepository) FindByPhone(ctx context.Context, phone string) (*domain.User, error) {
	const query = "SELECT id, name, enabled FROM %s WHERE phone = $1 LIMIT 1"

	user := &domain.User{
		Phone: phone,
	}

	err := r.db.QueryRowContext(ctx, r.table(query), phone).Scan(&user.ID, &user.Name, &user.Enabled)
	return user, err
}

func (r UserRepository) Update(ctx context.Context, user *domain.User) error {
	const query = "UPDATE %s SET name = $2, enabled = $3 WHERE id = $1"

	_, err := r.db.ExecContext(ctx, r.table(query), user.ID, user.Name, user.Enabled)

	return err
}

func (r UserRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
