package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"strings"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/errorsShare"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (s *Storage) ListProductsKaspiSearch(ctx context.Context, registrator_id int64, params dto.ProductsQueryRequest) ([]models.ProductsItemEntity, int, error) {
	op := "postgres.ListProducts"
	log := s.log.With("op", op)

	var products = []models.ProductsItemEntity{}
	sb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	qntDistinct := sb.Select(
		"qpr.product_id",
		//"qpr.sum",
		"qpr.product_create_date",
		"qpr.vid_modeli_id",
		"qpr.product_group_id",
		"qpr.create_date",
	).
		From("qnt_price_registry qpr").
		Where(squirrel.Eq{"qpr.registrator_id": registrator_id}).
		Distinct()

	if params.Skip != 0 {
		qntDistinct = qntDistinct.Offset(uint64(params.Skip))
	}
	if params.Take != 0 {
		qntDistinct = qntDistinct.Limit(uint64(params.Take))
	}
	if len(params.Size) != 0 {
		qntDistinct = qntDistinct.Where(squirrel.Eq{"size_id": params.Size})
		//qntDistinctCount = qntDistinctCount.Where(squirrel.Eq{"size_id": params.Size})
	}
	if len(params.Product_group) != 0 {
		qntDistinct = qntDistinct.Where(squirrel.Eq{"product_group_id": params.Product_group})
	}
	if len(params.Vid_modeli) != 0 {
		qntDistinct = qntDistinct.Where(squirrel.Eq{"vid_modeli_id": params.Vid_modeli})
	}
	if params.MinPrice != 0 {
		qntDistinct = qntDistinct.Where(squirrel.GtOrEq{"sum": params.MinPrice})
	}
	if params.MaxPrice != 0 {
		qntDistinct = qntDistinct.Where(squirrel.LtOrEq{"sum": params.MaxPrice})
	}
	if params.Search_name != "" {
		qntDistinct = qntDistinct.Where(squirrel.ILike{"product_name": "%" + params.Search_name + "%"})
	}
	if params.Sort != "" {
		sortSlice := strings.Split(params.Sort, "-")
		log.Debug("sort", slog.Any("sortSlice", sortSlice))
		// делаем костыль если product_create_date не подходит для сотировки
		if sortSlice[0] == "product_create_date" {
			sortSlice[0] = "product_id"
		}
		qntDistinct = qntDistinct.OrderBy(sortSlice[0] + " " + sortSlice[1])
	} else {
		qntDistinct = qntDistinct.OrderBy("product_id desc")
	}

	mainQuery := `    q.product_id,
    q.product_create_date,
    q.sum,
    q.product_group_id,
    pg.name_1c AS product_group_name,
    p.name AS product_name,
    p.artikul AS artikul,
    p.name AS name,
    p.description AS description,
	p.material_podoshva AS material_podoshva,
	p.material_up AS material_up,
	p.material_inside AS material_inside,
	p.sex AS sex,
    v.name_1c AS vid_modeli_name,
    q.vid_modeli_id,
    q.create_date,
	(select json_agg(jsonb_build_object('id', im.id, 'full_name', im.full_name,
	'name', im.name, 'active', im.active, 'main', im.main	
	)) from image_registry im where im.product_id = q.product_id) as image_registry,
    (
        SELECT jsonb_agg(jsonb_build_object('store_id', sub.store_id, 'size', sub.size_name_1c, 'sum', sub.sum, 'qnt', sub.qnt))
        FROM (
            SELECT size_name_1c, sum, qnt, array_agg(distinct store_id) as store_id
            FROM qnt_price_registry
            WHERE product_id = q.product_id and sum = q.sum
            group by size_name_1c, sum, qnt, store_id
        ) sub
    ) AS qnt_price`

	queryBuilder := sb.Select(mainQuery).FromSelect(qntDistinct, "q").
		LeftJoin(`product p ON q.product_id = p.id`).
		LeftJoin(`vid_modeli v ON q.vid_modeli_id = v.id`).
		LeftJoin(`product_group pg ON q.product_group_id = pg.id`)

	query2, args, err2 := queryBuilder.ToSql()
	if err2 != nil {
		log.Error(err2.Error())
		return products, 0, errorsShare.ErrInternalError.Error
	}
	//log.Debug(query2)

	err := pgxscan.Select(ctx, s.Db, &products, query2, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// если выкидывается ошибка нет строк, возвращаем пустой массив
			return products, 0, nil
		}
		log.Error(err.Error())
		return products, 0, errorsShare.ErrInternalError.Error
	}

	// считаем кол-во строк без пагинации
	var productsCount = []models.ProductsItemEntity{}
	qntDistinctCount := qntDistinct.RemoveLimit().RemoveOffset()
	queryCount, argsCount, errCount := qntDistinctCount.ToSql()
	if errCount != nil {
		log.Error(errCount.Error())
		return products, 0, errorsShare.ErrInternalError.Error
	}
	//log.Debug(queryCount)
	err = pgxscan.Select(ctx, s.Db, &productsCount, queryCount, argsCount...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// если выкидывается ошибка нет строк, возвращаем пустой массив
			return products, 0, nil
		}
		log.Error(err.Error())
		return products, 0, errorsShare.ErrInternalError.Error
	}

	// заполняем image_active_path
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
	fullCount := len(productsCount)

	return products, fullCount, nil
}
