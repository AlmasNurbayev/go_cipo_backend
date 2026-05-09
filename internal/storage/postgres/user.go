package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/errorsShare"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgconn"
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

func (s *Storage) GetUserByEmailStorage(ctx context.Context, email string) (models.UserEntity, error) {
	op := "postgres.GetUserByEmailStorage"
	log := s.log.With("op", op)

	var user = models.UserEntity{}

	query := `SELECT id, name, email, role, password FROM "user" WHERE email = $1`
	err := pgxscan.Get(ctx, s.Db, &user, query, email)
	if err != nil {
		log.Error(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return user, errorsShare.ErrUsernameOrPasswordIsWrong.Error
		}
		return user, errorsShare.ErrInternalError.Error
	}

	return user, nil
}

func (s *Storage) CreateUserStorage(ctx context.Context, user models.UserEntity) (int64, error) {
	op := "postgres.CreateUserStorage"
	log := s.log.With("op", op)

	query := `INSERT INTO "user" (name, email, role, password) VALUES ($1, $2, $3, $4) RETURNING id;`
	err := s.Db.QueryRow(ctx, query, user.Name, user.Email, user.Role, user.Password).Scan(
		&user.Id)
	if err != nil {
		var pgErr *pgconn.PgError
		log.Error("error: ", slog.String("err", err.Error()))
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return 0, errorsShare.ErrUserAlreadyExists.Error
		} else {
			return 0, errorsShare.ErrInternalError.Error
		}
	}

	return user.Id, nil
}
