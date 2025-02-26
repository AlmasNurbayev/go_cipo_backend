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

func (s *Storage) CreateImageRegistry(ctx context.Context, data models.ImageRegistryEntity) (int64, error) {
	op := "postgres.CreateImageRegistry"
	log := s.log.With(slog.String("op", op))

	query := `INSERT INTO image_registry
	(main, main_change_date, resolution, size, full_name, name, path, 
	operation_date, active, active_change_date, registrator_id, product_id) 
		VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) 
		RETURNING id;`
	db := *s.Tx

	err := db.QueryRow(ctx, query,
		data.Main, data.Main_change_date, data.Resolution, data.Size,
		data.Full_name, data.Name, data.Path, data.Operation_date, data.Active,
		data.Active_change_date, data.Registrator_id, data.Product_id).Scan(
		&data.Id)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return data.Id, err
	}
	return data.Id, nil
}

func (s *Storage) UpdateImageRegistryByName(ctx context.Context, data models.ImageRegistryEntity) error {
	op := "postgres.UpdateImageRegistry"
	log := s.log.With(slog.String("op", op))

	query := `UPDATE image_registry SET
	main = $1, main_change_date = $2, resolution = $3, size = $4, full_name = $5, name = $6, path = $7, 
	operation_date = $8, active = $9, active_change_date = $10, registrator_id = $11, product_id = $12 
		WHERE name = $13 RETURNING *;`
	db := *s.Tx

	err := pgxscan.Get(ctx, db, &data, query,
		data.Main, data.Main_change_date, data.Resolution, data.Size,
		data.Full_name, data.Name, data.Path, data.Operation_date, data.Active,
		data.Active_change_date, data.Registrator_id, data.Product_id, data.Name)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return err
	}
	return nil
}

func (s *Storage) ListImageRegistry(ctx context.Context) ([]models.ImageRegistryEntity, error) {
	op := "postgres.ListImageRegistry"
	log := s.log.With(slog.String("op", op))

	var images = []models.ImageRegistryEntity{}

	query := `SELECT id, main, main_change_date, resolution, size, full_name, name, path, 
	operation_date, active, active_change_date, registrator_id, product_id FROM image_registry;`
	db := *s.Tx

	err := pgxscan.Select(ctx, db, &images, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// если выкидывается ошибка нет строк, возвращаем пустой массив
			return images, nil
		}
		log.Error("error: ", slog.String("err", err.Error()))
		return images, err
	}
	return images, nil
}

func (s *Storage) GetImagesByProductId(ctx context.Context, id int64) ([]models.ImageRegistryEntity, error) {
	op := "postgres.GetImagesByProductId"
	log := s.log.With(slog.String("op", op))

	var productImages = []models.ImageRegistryEntity{}

	db := s.Db
	query := `SELECT id, resolution, full_name, name, path, size, operation_date, 
	main, main_change_date, active, active_change_date, product_id, registrator_id, 
	create_date, changed_date FROM image_registry WHERE product_id = $1;`
	err := pgxscan.Select(ctx, db, &productImages, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// если выкидывается ошибка нет строк, возвращаем пустой массив
			return productImages, nil
		}
		log.Error(err.Error())
		return productImages, errorsShare.ErrInternalError.Error
	}
	return productImages, nil
}
