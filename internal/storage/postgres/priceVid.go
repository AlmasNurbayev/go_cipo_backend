package postgres

import (
	"context"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (s *Storage) CreatePriceVid(ctx context.Context, data models.PriceVidEntity) (int64, error) {
	op := "postgres.CreatePriceVid"
	log := s.log.With(slog.String("op", op))

	query := `INSERT INTO price_vid
	(id_1c, name_1c, registrator_id, active, active_change_date) 
		VALUES 
		($1, $2, $3, $4, $5) 
		RETURNING id;`
	db := *s.Tx

	err := db.QueryRow(ctx, query,
		data.Id_1c,
		data.Name_1c,
		data.Registrator_id,
		data.Active,
		data.Active_change_date,
	).Scan(
		&data.Id)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return data.Id, err
	}
	return data.Id, nil
}

func (s *Storage) UpdatePriceVid(ctx context.Context, data models.PriceVidEntity) error {
	op := "postgres.UpdatePriceVid"
	log := s.log.With(slog.String("op", op))

	query := `UPDATE price_vid SET
	id_1c = $1, name_1c = $2, registrator_id = $3, active = $4, active_change_date = $5 
		WHERE id_1c = $6 RETURNING *;`
	db := *s.Tx

	err := pgxscan.Get(ctx, db, &data, query,
		data.Id_1c,
		data.Name_1c,
		data.Registrator_id, data.Active, data.Active_change_date, data.Id_1c)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return err
	}
	return nil
}

func (s *Storage) ListPriceVid(ctx context.Context) ([]models.PriceVidEntity, error) {
	op := "postgres.ListPriceVid"
	log := s.log.With(slog.String("op", op))

	query := `SELECT id, id_1c, name_1c, registrator_id, active, 
	active_change_date FROM price_vid;`
	db := *s.Tx

	var priceVids []models.PriceVidEntity
	err := pgxscan.Select(ctx, db, &priceVids, query)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return nil, err
	}
	return priceVids, nil
}
