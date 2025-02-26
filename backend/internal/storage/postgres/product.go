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

	// json, err2 := utils.PrintAsJSON(data)
	// if err2 != nil {
	// 	log.Error("error: ", slog.String("err", err2.Error()))
	// 	return err2
	// }
	// fmt.Println(string(*json))

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

func (s *Storage) GetProductById(ctx context.Context, id int64) (models.ProductEntity, error) {
	op := "postgres.GetProductById"
	log := s.log.With(slog.String("op", op))

	var product = models.ProductEntity{}

	db := s.Db
	query := `SELECT * FROM product WHERE id = $1;`
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
