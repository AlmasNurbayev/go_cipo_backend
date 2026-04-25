package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (s *Storage) ListBrend(ctx context.Context) ([]models.BrendEntity, error) {
	op := "postgres.ListBrend"
	log := s.log.With(slog.String("op", op))

	query := `SELECT id, id_1c, name_1c, registrator_id FROM brend;`

	var brends []models.BrendEntity
	var err error
	// если есть транзакция, используем ее, иначе стандартный пул
	if s.Tx != nil {
		db := *s.Tx
		err = pgxscan.Select(ctx, db, &brends, query)
	} else {
		db := s.Db
		err = pgxscan.Select(ctx, db, &brends, query)
	}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// если выкидывается ошибка нет строк, возвращаем пустой массив
			return brends, nil
		}
		log.Error("error: ", slog.String("err", err.Error()))
		return brends, err
	}
	return brends, nil
}

func (s *Storage) CreateBrend(ctx context.Context, data models.BrendEntity) (int64, error) {
	op := "postgres.CreateBrend"
	log := s.log.With(slog.String("op", op))

	query := `INSERT INTO brend
	(id_1c, name_1c, registrator_id) 
		VALUES 
		($1, $2, $3) 
		RETURNING id;`
	db := *s.Tx

	err := db.QueryRow(ctx, query,
		data.Id_1c,
		data.Name_1c,
		data.Registrator_id).Scan(
		&data.Id)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return data.Id, err
	}
	return data.Id, nil
}

func (s *Storage) UpdateBrend(ctx context.Context, data models.BrendEntity) error {
	op := "postgres.UpdateBrend"
	log := s.log.With(slog.String("op", op))

	query := `UPDATE brend SET
	id_1c = $1, name_1c = $2, registrator_id = $3 
		WHERE id_1c = $4 RETURNING *;`
	db := *s.Tx

	err := pgxscan.Get(ctx, db, &data, query,
		data.Id_1c,
		data.Name_1c,
		data.Registrator_id, data.Id_1c)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return err
	}
	return nil
}
