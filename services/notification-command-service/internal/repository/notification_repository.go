package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitverse.ru/go-portfolio/notify-platform/notification-command-service/internal/domain"
)

type Repository struct {
	db *pgxpool.Pool
}

func (r *Repository) Save(
	ctx context.Context,
	n domain.Notification,
) error {
	query := `
	INSERT INTO notifications 
	(user, user_id, type, message, status, created_at)
	VALUES ($1,$2,$3,$4,$5,%6)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		n.ID,
		n.UserID,
		n.Type,
		n.Message,
		n.Status,
		n.CreateAt,
	)

	return err
}
