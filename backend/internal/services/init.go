package services

import (
	"context"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/config"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
)

type Service struct {
	log             *slog.Logger
	postgresStorage postgresStorage
	cfg             *config.Config
}

type postgresStorage interface {
	GetUserByIdStorage(ctx context.Context, id int64) (models.UserEntity, error)
	GetUserByNameStorage(ctx context.Context, name string) (models.UserEntity, error)

	ListSize(ctx context.Context) ([]models.SizeEntity, error)
	ListStore(ctx context.Context) ([]models.StoreEntity, error)
	ListVidModeli(ctx context.Context) ([]models.VidModeliEntity, error)
	ListProductGroup(ctx context.Context) ([]models.ProductsGroupEntity, error)
	ListBrend(ctx context.Context) ([]models.BrendEntity, error)

	GetProductById(ctx context.Context, id int64) (models.ProductEntity, error)
	GetImagesByProductId(ctx context.Context, id int64) ([]models.ImageRegistryEntity, error)

	GetProductGroupById(ctx context.Context, id int64) (models.ProductsGroupEntity, error)
	GetVidModeliById(ctx context.Context, id int64) (models.VidModeliEntity, error)
	GetLastOfferRegistrator(ctx context.Context) (models.RegistratorEntity, error)

	GetQntPriceRegistryByProductId(ctx context.Context, product_id int64, registrator_id int64) ([]models.QntPriceRegistryEntity, error)
}

func NewService(log *slog.Logger,
	postgresStorage postgresStorage,
	cfg *config.Config) *Service {
	return &Service{
		log:             log,
		postgresStorage: postgresStorage,
		cfg:             cfg,
	}
}
