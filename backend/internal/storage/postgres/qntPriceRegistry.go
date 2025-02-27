package postgres

import (
	"context"
	"log/slog"
	"time"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (s *Storage) CreateQntPriceRegistry(ctx context.Context, data models.QntPriceRegistryEntity) (int64, error) {
	op := "postgres.CreateQntPriceRegistry"
	log := s.log.With(slog.String("op", op))

	query := `INSERT INTO qnt_price_registry
	(registrator_id, sum, qnt, operation_date, discount_percent, 
	discount_begin, discount_end, store_id, product_id, price_vid_id, size_id, 
	product_group_id, vid_modeli_id, size_name_1c, product_name, product_create_date) 
		VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) 
		RETURNING id;`
	db := *s.Tx

	err := db.QueryRow(ctx, query,
		data.Registrator_id, data.Sum, data.Qnt, data.Operation_date,
		data.Discount_percent, data.Discount_begin, data.Discount_end,
		data.Store_id, data.Product_id, data.Price_vid_id, data.Size_id,
		data.Product_group_id, data.Vid_modeli_id, data.Size_name_1c,
		data.Product_name, data.Product_create_date).Scan(
		&data.Id)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return data.Id, err
	}
	return data.Id, nil
}

func (s *Storage) SearchQntPriceRegistry(ctx context.Context, product_id int64, registrator_id int64) ([]models.QntPriceRegistryEntity, error) {
	op := "postgres.SearchQntPriceRegistry"
	log := s.log.With(slog.String("op", op))

	query := `SELECT * FROM qnt_price_registry WHERE product_id = $1 AND registrator_id = $2;`
	db := *s.Tx

	var qntPriceRegistry []models.QntPriceRegistryEntity
	err := pgxscan.Select(ctx, db, &qntPriceRegistry, query, product_id, registrator_id)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return nil, err
	}
	return qntPriceRegistry, nil
}

func (s *Storage) GetQntPriceRegistryBeforeDate(ctx context.Context, date time.Time) ([]models.QntPriceRegistryEntity, error) {
	op := "postgres.GetQntPriceRegistryBeforeDate"
	log := s.log.With(slog.String("op", op))

	query := `SELECT id FROM qnt_price_registry WHERE create_date < $1;`
	db := *s.Tx

	var qntPriceRegistry []models.QntPriceRegistryEntity
	err := pgxscan.Select(ctx, db, &qntPriceRegistry, query, date)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return nil, err
	}
	return qntPriceRegistry, nil
}

func (s *Storage) DeleteQntPriceRegistryById(ctx context.Context, ids []int64) error {
	op := "postgres.DeleteQntPriceRegistryById"
	log := s.log.With(slog.String("op", op))

	query := `DELETE FROM qnt_price_registry WHERE id = ANY($1);`
	db := *s.Tx

	_, err := db.Exec(ctx, query, ids)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return nil
	}
	return nil
}

func (s *Storage) GetQntPriceRegistryByProductId(ctx context.Context, product_id int64, registrator_id int64) ([]models.QntPriceRegistryEntityByProduct, error) {
	op := "postgres.GetQntPriceRegistryByProductId"
	log := s.log.With(slog.String("op", op))

	query := `SELECT size_id, size_name_1c, qnt, sum, store_id 
	FROM qnt_price_registry 
	WHERE product_id = $1 and registrator_id = $2;`
	db := s.Db

	var qntPriceRegistry []models.QntPriceRegistryEntityByProduct
	err := pgxscan.Select(ctx, db, &qntPriceRegistry, query, product_id, registrator_id)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return nil, err
	}
	return qntPriceRegistry, nil
}

func (s *Storage) GetQntPriceRegistryGroupByProductId(ctx context.Context, product_id int64, registrator_id int64) ([]models.QntPriceRegistryEntityGroupByProduct, error) {
	op := "postgres.GetQntPriceRegistryGroupByProductId"
	log := s.log.With(slog.String("op", op))

	//json_agg (DISTINCT store_id) as store_id
	query := `select
  size_id,
  size_name_1c,
  SUM(sum) AS sum,
  SUM(qnt) AS qnt,
	json_agg (DISTINCT store_id) as store_id
	FROM qnt_price_registry
	WHERE product_id = $1 and registrator_id = $2
	GROUP BY size_name_1c, size_id;`
	db := s.Db

	var res []models.QntPriceRegistryEntityGroupByProduct
	err := pgxscan.Select(ctx, db, &res, query, product_id, registrator_id)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return nil, err
	}
	return res, nil
}
