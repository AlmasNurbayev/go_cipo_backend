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

	var excludeIds = []int64{}
	//fmt.Println(excludeIds)
	if len(s.Cfg.HTTP.EXCLUDE_VIDS_IN_LIST) > 0 {
		var err error
		excludeIds, err = s.ListVidModeliIdExcludeNames(ctx, s.Cfg.HTTP.EXCLUDE_VIDS_IN_LIST)
		if err != nil {
			return products, err
		}
	}

	db := s.Db

	query := `SELECT p.id as "product_id", p.name as "product_name",
	p.create_date as "product_create_date", p.name as "name", artikul as "artikul",
	p.description as "description", p.material_up as "material_up", 
	p.material_inside as "material_inside", p.material_podoshva as "material_podoshva", 
	p.sex as "sex", 
	p.vid_modeli_id as "vid_modeli_id",
	p.nom_vid as "nom_vid",
	vid.name_1c as "vid_modeli_name", 
	pg.name_1c as "product_group_name",
	imr.full_name as "image_active_path",
	(select json_agg(jsonb_build_object('size', agg.size_name_1c, 'sum', agg.sum)) FROM (
		SELECT 
		iq.size_name_1c,
		SUM(iq.sum) AS sum
		FROM qnt_price_registry iq
		WHERE iq.product_id = p.id and iq.registrator_id = $1
		GROUP BY iq.size_name_1c
	) agg) as qnt_price  
	FROM product p
	LEFT JOIN vid_modeli vid ON vid_modeli_id = vid.id
	LEFT JOIN product_group pg ON product_group_id = pg.id
	LEFT JOIN ( SELECT DISTINCT ON (product_id) *
	    FROM image_registry
	    ORDER BY product_id, main DESC, id) imr ON imr.product_id = p.id
	WHERE 
		vid.id <> ALL($3)
		AND EXISTS (
				SELECT 1
				FROM qnt_price_registry iq
				WHERE iq.qnt > 0
					AND iq.product_id = p.id
					AND iq.registrator_id = $1)
    AND imr.full_name <> ''
	ORDER BY product_id desc
	LIMIT $2;`

	err := pgxscan.Select(ctx, db, &products, query, registrator_id, count, excludeIds)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// если выкидывается ошибка нет строк, возвращаем пустой массив
			return products, nil
		}
		log.Error(err.Error())
		return products, errorsShare.ErrInternalError.Error
	}

	return products, nil

}
