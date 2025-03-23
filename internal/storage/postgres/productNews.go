package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/errorsShare"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (s *Storage) ListProductNews(ctx context.Context, registrator_id int64, count int64) ([]models.ProductNewsEntity, error) {
	op := "postgres.ListProductNews"
	log := s.log.With("op", op)

	var products = []models.ProductNewsEntity{}
	db := s.Db

	query := `SELECT p.id as "product_id", p.name as "product_name",
	p.create_date as "product_create_date", p.name as "name", artikul as "artikul",
	p.description as "description", p.material_up as "material_up", 
	p.material_inside as "material_inside", p.material_podoshva as "material_podoshva", 
	p.sex as "sex", 
	p.vid_modeli_id as "vid_modeli_id",
	vid.name_1c as "vid_modeli_name", 
	pg.name_1c as "product_group_name",
	qpr.sum as "sum",
	(select json_agg(jsonb_build_object('id', im.id, 'full_name', im.full_name,
	'name', im.name, 'active', im.active, 'main', im.main	
	)) from image_registry im where im.product_id = p.id) as image_registry,
	(select json_agg(jsonb_build_object('size', agg.size_name_1c, 'sum', agg.sum,
	'qnt', agg.qnt, 'store_id', agg.store_id)) FROM (
		SELECT 
		json_agg (distinct iq.store_id) as store_id,
		iq.size_name_1c,
		SUM(iq.sum) AS sum,
		SUM(iq.qnt) AS qnt
		FROM qnt_price_registry iq
		WHERE iq.product_id = p.id
		GROUP BY iq.store_id, iq.size_name_1c
	) agg) as qnt_price  
	FROM product p
	LEFT JOIN vid_modeli vid ON vid_modeli_id = vid.id
	LEFT JOIN product_group pg ON product_group_id = pg.id
	JOIN (
    SELECT product_id, sum
    FROM qnt_price_registry 
    WHERE registrator_id = $1 
    GROUP BY sum, product_id 
    ORDER BY product_id DESC 
    LIMIT $2
	) qpr ON qpr.product_id = p.id
	ORDER BY product_id desc;`

	err := pgxscan.Select(ctx, db, &products, query, registrator_id, count)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// если выкидывается ошибка нет строк, возвращаем пустой массив
			return products, nil
		}
		log.Error(err.Error())
		return products, errorsShare.ErrInternalError.Error
	}

	for index, value := range products {
		if value.Image_active_path != "" {
			continue
		}
		for _, image := range value.Image_registry {
			if image.Main {
				products[index].Image_active_path = image.Full_name
				break
			}
		}
	}

	return products, nil

}
