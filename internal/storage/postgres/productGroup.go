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

func (s *Storage) CreateProductGroup(ctx context.Context, data models.ProductsGroupEntity) (int64, error) {
	op := "postgres.CreateProductGroup"
	log := s.log.With(slog.String("op", op))

	query := `INSERT INTO product_group
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

func (s *Storage) ListProductGroup(ctx context.Context) ([]models.ProductsGroupEntity, error) {
	op := "postgres.ListProductsGroup"
	log := s.log.With("op", op)

	var productsGroups = []models.ProductsGroupEntity{}

	query := `SELECT id, id_1c, name_1c, registrator_id FROM product_group;`

	var err error
	// если есть транзакция, используем ее, иначе стандартный пул
	if s.Tx != nil {
		db := *s.Tx
		err = pgxscan.Select(ctx, db, &productsGroups, query)
	} else {
		db := s.Db
		err = pgxscan.Select(ctx, db, &productsGroups, query)
	}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// если выкидывается ошибка нет строк, возвращаем пустой массив
			return productsGroups, nil
		}
		log.Error(err.Error())
		return productsGroups, errorsShare.ErrInternalError.Error
	}
	return productsGroups, nil
}

func (s *Storage) UpdateProductGroup(ctx context.Context, data models.ProductsGroupEntity) error {
	op := "postgres.UpdateProductGroup"
	log := s.log.With(slog.String("op", op))

	query := `UPDATE product_group SET
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

func (s *Storage) GetProductGroupById(ctx context.Context, id int64) (models.ProductsGroupEntity, error) {
	op := "postgres.GetProductGroupById"
	log := s.log.With(slog.String("op", op))

	query := `SELECT * FROM product_group WHERE id = $1;`
	db := s.Db

	data := models.ProductsGroupEntity{}

	err := pgxscan.Get(ctx, db, &data, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return data, errorsShare.ErrInternalError.Error
		}
		log.Error("error: ", slog.String("err", err.Error()))
		return data, err
	}
	return data, nil
}
