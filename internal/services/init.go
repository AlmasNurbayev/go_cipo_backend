package services

import (
	"context"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/config"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
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

	GetProductById(ctx context.Context, id int64) (models.ProductByIdEntity, error)
	GetProductByName1c(ctx context.Context, name_1c string) (models.ProductByIdEntity, error)
	GetImagesByProductId(ctx context.Context, id int64) ([]models.ImageRegistryEntity, error)

	GetProductGroupById(ctx context.Context, id int64) (models.ProductsGroupEntity, error)
	GetVidModeliById(ctx context.Context, id int64) (models.VidModeliEntity, error)
	GetLastOfferRegistrator(ctx context.Context) (models.RegistratorEntity, error)

	GetQntPriceRegistryByProductId(ctx context.Context, product_id int64, registrator_id int64) ([]models.QntPriceRegistryEntityByProduct, error)
	GetQntPriceRegistryGroupByProductId(ctx context.Context, product_id int64, registrator_id int64) ([]models.QntPriceRegistryEntityGroupByProduct, error)

	ListProductNews(ctx context.Context, registrator_id int64, count int64) ([]models.ProductNewsEntity, error)

	GetNewsById(ctx context.Context, id int64) (models.NewsEntity, error)
	ListNews(ctx context.Context, count int64) ([]models.NewsEntity, error)

	ListProductsSearch(ctx context.Context, registrator_id int64, params dto.ProductsQueryRequest) ([]models.ProductsItemEntity, int, error)

	CreateKaspiOrganization(ctx context.Context, data models.KaspiOrganizationEntity) (models.KaspiOrganizationEntity, error)
	ListKaspiOrganization(ctx context.Context) ([]models.KaspiOrganizationEntity, error)
	CreateKaspiCategory(ctx context.Context, data models.KaspiCategoriesEntity) (models.KaspiCategoriesEntity, error)
	ListKaspiCategory(ctx context.Context) ([]models.KaspiCategoriesEntity, error)

	ListKaspiProductsSearch(ctx context.Context, registrator_id int64, params dto.KaspiProductsQueryRequest) ([]models.ProductsItemEntity, int, error)
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
