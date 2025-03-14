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

func (s *Storage) ListProduct(ctx context.Context) ([]models.ProductEntity, error) {
	op := "postgres.ListProduct"
	log := s.log.With("op", op)

	var products = []models.ProductEntity{}

	query := `SELECT * FROM product;`
	err := pgxscan.Select(ctx, *s.Tx, &products, query)
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

func (s *Storage) CreateProduct(ctx context.Context, data models.ProductEntity) (int64, error) {
	op := "postgres.CreateProduct"
	log := s.log.With(slog.String("op", op))

	query := `INSERT INTO product
	(id_1c, name_1c, registrator_id, name, product_group_id, product_vid_id,
		 vid_modeli_id, artikul, base_ed, description, material_inside, 
		 material_podoshva, material_up, sex, product_folder, main_color, public_web) 
		VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) 
		RETURNING id;`
	db := *s.Tx

	err := db.QueryRow(ctx, query,
		data.Id_1c, data.Name_1c, data.Registrator_id, data.Name, data.Product_group_id,
		data.Product_vid_id, data.Vid_modeli_id, data.Artikul, data.Base_ed, data.Description,
		data.Material_inside, data.Material_podoshva, data.Material_up, data.Sex,
		data.Product_folder, data.Main_color, data.Public_web).Scan(
		&data.Id)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return data.Id, err
	}
	return data.Id, nil
}

func (s *Storage) UpdateProductById1c(ctx context.Context, data models.ProductEntity) error {
	op := "postgres.UpdateProduct"
	log := s.log.With(slog.String("op", op))

	query := `UPDATE product SET
	id_1c = $1, name_1c = $2, registrator_id = $3, name = $4, 
	product_group_id = $5, product_vid_id = $6, vid_modeli_id = $7, 
	artikul = $8, base_ed = $9, description = $10, material_inside = $11, 
	material_podoshva = $12, material_up = $13, sex = $14, 
	product_folder = $15, main_color = $16, public_web = $17
		WHERE id_1c = $18 RETURNING *;`
	db := *s.Tx

	err := pgxscan.Get(ctx, db, &data, query,
		data.Id_1c, data.Name_1c, data.Registrator_id, data.Name, data.Product_group_id,
		data.Product_vid_id, data.Vid_modeli_id, data.Artikul, data.Base_ed, data.Description,
		data.Material_inside, data.Material_podoshva, data.Material_up, data.Sex,
		data.Product_folder, data.Main_color, data.Public_web, data.Id_1c)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return err
	}
	return nil
}

func (s *Storage) GetProductById(ctx context.Context, id int64) (models.ProductByIdEntity, error) {
	op := "postgres.GetProductById"
	log := s.log.With(slog.String("op", op))

	var product = models.ProductByIdEntity{}

	db := s.Db
	query := `SELECT p.*, 
	vid.id as "vid_modeli.id", vid.name_1c as "vid_modeli.name_1c", 
	pg.id as "product_group.id", pg.name_1c as "product_group.name_1c" 
	FROM product p
	LEFT JOIN vid_modeli vid ON p.vid_modeli_id = vid.id
	LEFT JOIN product_group pg ON p.product_group_id = pg.id
	WHERE p.id = $1;`
	err := pgxscan.Get(ctx, db, &product, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// если выкидывается ошибка нет строк, возвращаем пустой массив
			return product, errorsShare.ErrProductNotFound.Error
		}
		log.Error(err.Error())
		return product, errorsShare.ErrInternalError.Error
	}
	return product, nil
}

func (s *Storage) ListProductsSearch(ctx context.Context, registrator_id int64, params dto.ProductsQueryRequest) ([]models.ProductsItemEntity, int, error) {
	op := "postgres.ListProducts"
	log := s.log.With("op", op)

	var products = []models.ProductsItemEntity{}
	sb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	qntDistinct := sb.Select(
		"qpr.product_id",
		"qpr.sum",
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

	log.Info("Count query")

	// считаем кол-во строк без пагинации
	var productsCount = []models.ProductsItemEntity{}
	qntDistinctCount := qntDistinct.RemoveLimit().RemoveOffset()
	queryCount, argsCount, errCount := qntDistinctCount.ToSql()
	if errCount != nil {
		log.Error(errCount.Error())
		return products, 0, errorsShare.ErrInternalError.Error
	}
	log.Debug(queryCount)
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
