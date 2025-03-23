package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/errorsShare"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (s *Storage) GetUserByIdStorage(ctx context.Context, id int64) (models.UserEntity, error) {
	op := "postgres.GetUserByIdStorage"
	log := s.log.With("op", op)

	// искусственное замедление запроса
	// var temp string
	// err2 := pgxscan.Get(ctx, s.db, &temp, "SELECT pg_sleep(18)")
	// if err2 != nil {
	// 	s.log.Error("canceled query DB", "error", err2)
	// 	return user, err2
	// }

	var user = models.UserEntity{}

	query := `SELECT id, name, email, role FROM "user" WHERE id = $1`
	err := pgxscan.Get(ctx, s.Db, &user, query, id)
	if err != nil {
		log.Error(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return user, errorsShare.ErrUserNotFound.Error
		}
		return user, errorsShare.ErrInternalError.Error
	}

	return user, nil
}

func (s *Storage) GetUserByNameStorage(ctx context.Context, name string) (models.UserEntity, error) {
	op := "postgres.GetUserByNameStorage"
	log := s.log.With("op", op)

	var user = models.UserEntity{}

	query := `SELECT id, name, email, role FROM "user" WHERE name = $1`
	err := pgxscan.Get(ctx, s.Db, &user, query, name)
	if err != nil {
		log.Error(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return user, errorsShare.ErrUserNotFound.Error
		}
		return user, errorsShare.ErrInternalError.Error
	}

	return user, nil
}
