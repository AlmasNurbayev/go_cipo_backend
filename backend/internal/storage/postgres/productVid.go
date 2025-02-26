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

func (s *Storage) CreateProductVid(ctx context.Context, data models.ProductVidEntity) (int64, error) {
	op := "postgres.CreateProductVid"
	log := s.log.With(slog.String("op", op))

	query := `INSERT INTO product_vid
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

func (s *Storage) ListProductVid(ctx context.Context) ([]models.ProductVidEntity, error) {
	op := "postgres.ListProductsVid"
	log := s.log.With("op", op)

	var productstVid = []models.ProductVidEntity{}

	query := `SELECT id, id_1c, name_1c, registrator_id FROM product_vid;`
	err := pgxscan.Select(ctx, *s.Tx, &productstVid, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// если выкидывается ошибка нет строк, возвращаем пустой массив
			return productstVid, nil
		}
		log.Error(err.Error())
		return productstVid, errorsShare.ErrInternalError.Error
	}
	return productstVid, nil
}

func (s *Storage) UpdateProductVid(ctx context.Context, data models.ProductVidEntity) error {
	op := "postgres.UpdateProductVid"
	log := s.log.With(slog.String("op", op))

	query := `UPDATE product_vid SET
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
