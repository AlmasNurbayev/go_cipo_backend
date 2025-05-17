package parserService

import (
	"context"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/config"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
)

type storageProvider interface {
	CreateRegistrator(ctx context.Context, data models.RegistratorEntity) (int64, error)
	GetRegistratorById(ctx context.Context, id int64) (models.RegistratorEntity, error)
	CreateProductGroup(ctx context.Context, data models.ProductsGroupEntity) (int64, error)
	ListProductGroup(ctx context.Context) ([]models.ProductsGroupEntity, error)
	UpdateProductGroup(ctx context.Context, data models.ProductsGroupEntity) error
	CreateProductVid(ctx context.Context, data models.ProductVidEntity) (int64, error)
	ListProductVid(ctx context.Context) ([]models.ProductVidEntity, error)
	UpdateProductVid(ctx context.Context, data models.ProductVidEntity) error
	CreateVidModeli(ctx context.Context, data models.VidModeliEntity) (int64, error)
	ListVidModeli(ctx context.Context) ([]models.VidModeliEntity, error)
	UpdateVidModeli(ctx context.Context, data models.VidModeliEntity) error
	ListProductDescMapping(ctx context.Context) ([]models.ProductsDescMappingEntity, error)
	ListProduct(ctx context.Context) ([]models.ProductEntity, error)
	CreateProduct(ctx context.Context, data models.ProductEntity) (int64, error)
	UpdateProductById1c(ctx context.Context, data models.ProductEntity) error
	CreateImageRegistry(ctx context.Context, data models.ImageRegistryEntity) (int64, error)
	ListImageRegistry(ctx context.Context) ([]models.ImageRegistryEntity, error)
	UpdateImageRegistryByName(ctx context.Context, data models.ImageRegistryEntity) error
	CreateSize(ctx context.Context, data models.SizeEntity) (int64, error)
	ListSize(ctx context.Context) ([]models.SizeEntity, error)
	UpdateSize(ctx context.Context, data models.SizeEntity) error
	CreatePriceVid(ctx context.Context, data models.PriceVidEntity) (int64, error)
	ListPriceVid(ctx context.Context) ([]models.PriceVidEntity, error)
	UpdatePriceVid(ctx context.Context, data models.PriceVidEntity) error
	CreateStore(ctx context.Context, data models.StoreEntity) (int64, error)
	ListStore(ctx context.Context) ([]models.StoreEntity, error)
	UpdateStoreFrom1C(ctx context.Context, data models.StoreEntity) error
	CreateQntPriceRegistry(ctx context.Context, data models.QntPriceRegistryEntity) (int64, error)
}

type ParserService struct {
	ctx     context.Context
	storage storageProvider
	log     *slog.Logger
	cfg     *config.Config
}

func NewParserService(ctx context.Context, storage storageProvider, log *slog.Logger, cfg *config.Config) *ParserService {
	return &ParserService{
		ctx:     ctx,
		storage: storage,
		log:     log,
		cfg:     cfg,
	}
}
