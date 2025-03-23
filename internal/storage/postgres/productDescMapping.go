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

func (s *Storage) ListProductDescMapping(ctx context.Context) ([]models.ProductsDescMappingEntity, error) {
	op := "postgres.ListProductDescMapping"
	log := s.log.With(slog.String("op", op))

	var productDescMappingList = []models.ProductsDescMappingEntity{}

	query := `SELECT id, id_1c, name_1c, field FROM product_desc_mapping;`
	err := pgxscan.Select(ctx, *s.Tx, &productDescMappingList, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// если выкидывается ошибка нет строк, возвращаем пустой массив
			return productDescMappingList, nil
		}
		log.Error(err.Error())
		return productDescMappingList, errorsShare.ErrInternalError.Error
	}
	return productDescMappingList, nil
}
