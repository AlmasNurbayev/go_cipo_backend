package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (s *Storage) CreateKaspiOrganization(ctx context.Context, data models.KaspiOrganizationEntity) (models.KaspiOrganizationEntity, error) {
	op := "postgres.CreateOrganization"
	log := s.log.With(slog.String("op", op))

	query := `INSERT INTO kaspi_organizations
	(name, kaspi_id, kaspi_api_token_hash) 
		VALUES 
		($1, $2, $3) 
		RETURNING *;`
	db := s.Db

	result := models.KaspiOrganizationEntity{}

	err := pgxscan.Get(ctx, db, &result, query,
		data.Name,
		data.Kaspi_id,
		data.Kaspi_api_token_hash,
	)

	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return result, err
	}
	return result, nil
}

func (s *Storage) ListKaspiOrganization(ctx context.Context) ([]models.KaspiOrganizationEntity, error) {
	op := "postgres.ListOrganization"
	log := s.log.With(slog.String("op", op))

	query := `SELECT id, name, kaspi_id, kaspi_api_token_hash, is_active, create_date, changed_date FROM kaspi_organizations ORDER BY id DESC;`

	var result []models.KaspiOrganizationEntity

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
