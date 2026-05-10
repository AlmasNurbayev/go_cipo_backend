package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/errorsShare"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/gofiber/fiber/v3/log"
)

func (s *Storage) CreateKaspiExportGoodsRegistry(ctx context.Context, data models.KaspiExportGoodsRegistryEntity) (int64, error) {
	var id int64
	query := `INSERT INTO kaspi_export_goods_registry 
	(kaspi_organization_id, product_id, sended_body, 
	sended_category, sended_status) 
	VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := s.Db.QueryRow(ctx, query, data.KaspiOrganizationId,
		data.ProductId, data.SendedBody, data.SendedCategory,
		data.SendedStatus).Scan(&id)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return data.Id, err
	}
	return data.Id, nil
}

func (s *Storage) GetKaspiExportGoodsRegistryByProductId(ctx context.Context, productId int64) (models.KaspiExportGoodsRegistryEntity, error) {
	var data models.KaspiExportGoodsRegistryEntity
	query := `SELECT id, kaspi_organization_id, product_id, sended_body, 
	sended_category, sended_status, response_id, response_status, created_date, changed_date 
	FROM kaspi_export_goods_registry WHERE product_id = $1`
	err := s.Db.QueryRow(ctx, query, productId).Scan(&data.Id, &data.KaspiOrganizationId, &data.ProductId, &data.SendedBody, &data.SendedCategory, &data.SendedStatus, &data.ResponseId, &data.ResponseStatus, &data.CreatedDate, &data.ChangedDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// если выкидывается ошибка нет строк, возвращаем пустой массив
			return data, errorsShare.ErrKaspiExportGoodsRegistryItemNotFound.Error
		}
		log.Error("error: ", slog.String("err", err.Error()))
		return data, err
	}
	return data, nil
}

func (s *Storage) ListKaspiExportGoodsRegistry(ctx context.Context) ([]models.KaspiExportGoodsRegistryEntity, error) {
	var data []models.KaspiExportGoodsRegistryEntity
	query := `SELECT id, kaspi_organization_id, product_id, sended_body, 
	sended_category, sended_status, response_id, response_status, created_date, changed_date 
	FROM kaspi_export_goods_registry`
	err := s.Db.QueryRow(ctx, query).Scan(&data)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return data, err
	}
	return data, nil
}
