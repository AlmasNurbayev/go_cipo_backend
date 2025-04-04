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

func (s *Storage) CreateKaspiCategory(ctx context.Context, data models.KaspiCategoriesEntity) (models.KaspiCategoriesEntity, error) {
	op := "postgres.CreateKaspiCategory"
	log := s.log.With(slog.String("op", op))

	query := `INSERT INTO kaspi_categories
	(first_size, last_size, size_kaspi, name_kaspi, title_kaspi, gender_kaspi, 
	model_kaspi, material_kaspi, season_kaspi, colour_kaspi, attributes_list) 
		VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) 
		RETURNING *;`
	db := s.Db

	result := models.KaspiCategoriesEntity{}

	err := pgxscan.Get(ctx, db, &result, query,
		data.First_size,
		data.Last_size,
		data.Size_kaspi,
		data.Name_kaspi,
		data.Title_kaspi,
		data.Gender_kaspi,
		data.Model_kaspi,
		data.Material_kaspi,
		data.Season_kaspi,
		data.Colour_kaspi,
		data.Attributes_list,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		log.Error("error: ", slog.String("err", err.Error()))
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return result, errorsShare.ErrKaspiCategoryDuplicate.Error
		} else {
			return result, errorsShare.ErrInternalError.Error
		}
	}
	return result, nil
}

func (s *Storage) ListKaspiCategory(ctx context.Context) ([]models.KaspiCategoriesEntity, error) {
	op := "postgres.ListKaspiCategory"
	log := s.log.With(slog.String("op", op))

	query := `SELECT 
		id,
		first_size,
		last_size,
		size_kaspi,
		name_kaspi,
		title_kaspi,
		gender_kaspi,
		model_kaspi,
		material_kaspi,
		season_kaspi,
		colour_kaspi,
		attributes_list
	FROM kaspi_categories ORDER BY id DESC;`

	var result []models.KaspiCategoriesEntity

	db := s.Db
	err := pgxscan.Select(ctx, db, &result, query)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// если выкидывается ошибка нет строк, возвращаем пустой массив
			return result, nil
		}
		log.Error("error: ", slog.String("err", err.Error()))
		return result, err
	}
	return result, nil

}
