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
