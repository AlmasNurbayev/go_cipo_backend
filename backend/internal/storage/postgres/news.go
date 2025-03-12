package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/errorsShare"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (s *Storage) GetNewsById(ctx context.Context, id int64) (models.NewsEntity, error) {
	op := "postgres.GetNewsById"
	log := s.log.With("op", op)

	var news = models.NewsEntity{}

	query := `SELECT id, title, data, image_path, operation_date, changed_date FROM news WHERE id = $1`
	err := pgxscan.Get(ctx, s.Db, &news, query, id)
	if err != nil {
		log.Error(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return news, errorsShare.ErrNewsNotFound.Error
		}
		return news, errorsShare.ErrNewsNotFound.Error
	}
	return news, nil
}

func (s *Storage) ListNews(ctx context.Context, count int64) ([]models.NewsEntity, error) {
	op := "postgres.ListNews"
	log := s.log.With(slog.String("op", op))

	query := `SELECT id, title, data, image_path, operation_date, changed_date FROM news ORDER BY operation_date DESC LIMIT $1;`

	var news []models.NewsEntity
	var err error
	// если есть транзакция, используем ее, иначе стандартный пул

	db := s.Db
	err = pgxscan.Select(ctx, db, &news, query, count)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// если выкидывается ошибка нет строк, возвращаем пустой массив
			return news, nil
		}
		log.Error("error: ", slog.String("err", err.Error()))
		return news, err
	}
	return news, nil
}
