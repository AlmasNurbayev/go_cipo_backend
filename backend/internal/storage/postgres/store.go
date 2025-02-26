package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (s *Storage) CreateStore(ctx context.Context, data models.StoreEntity) (int64, error) {
	op := "postgres.CreateStore"
	log := s.log.With(slog.String("op", op))

	query := `INSERT INTO store
	(id_1c, name_1c, registrator_id, address, link_2gis, phone, 
	city, image_path, public, working_hours, yandex_widget_url, store_kaspi_id, 
	doublegis_widget_url) 
		VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) 
		RETURNING id;`
	db := *s.Tx

	err := db.QueryRow(ctx, query,
		data.Id_1c,
		data.Name_1c,
		data.Registrator_id, data.Address, data.Link_2gis,
		data.Phone, data.City, data.Image_path, data.Public, data.Working_hours,
		data.Yandex_widget_url, data.Store_kaspi_id, data.Doublegis_widget_url).Scan(
		&data.Id)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return data.Id, err
	}
	return data.Id, nil
}

func (s *Storage) UpdateStore(ctx context.Context, data models.StoreEntity) error {
	op := "postgres.UpdateStore"
	log := s.log.With(slog.String("op", op))

	query := `UPDATE store SET
	id_1c = $1, name_1c = $2, registrator_id = $3, address = $4, link_2gis = $5, 
	phone = $6, city = $7, image_path = $8, public = $9, working_hours = $10, 
	yandex_widget_url = $11, store_kaspi_id = $12, doublegis_widget_url = $13 
		WHERE id_1c = $14 RETURNING *;`
	db := *s.Tx

	err := pgxscan.Get(ctx, db, &data, query,
		data.Id_1c,
		data.Name_1c,
		data.Registrator_id, data.Address, data.Link_2gis,
		data.Phone, data.City, data.Image_path, data.Public, data.Working_hours,
		data.Yandex_widget_url, data.Store_kaspi_id, data.Doublegis_widget_url, data.Id_1c)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return err
	}
	return nil
}

func (s *Storage) ListStore(ctx context.Context) ([]models.StoreEntity, error) {
	op := "postgres.ListStore"
	log := s.log.With(slog.String("op", op))

	query := `SELECT id, id_1c, name_1c, registrator_id, address, link_2gis, 
	phone, city, image_path, public, working_hours, yandex_widget_url, store_kaspi_id, 
	doublegis_widget_url FROM store;`

	var stores []models.StoreEntity
	var err error
	// если есть транзакция, используем ее, иначе стандартный пул
	if s.Tx != nil {
		db := *s.Tx
		err = pgxscan.Select(ctx, db, &stores, query)
	} else {
		db := s.Db
		err = pgxscan.Select(ctx, db, &stores, query)
	}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// если выкидывается ошибка нет строк, возвращаем пустой массив
			return stores, nil
		}
		log.Error("error: ", slog.String("err", err.Error()))
		return stores, err
	}
	return stores, nil
}
