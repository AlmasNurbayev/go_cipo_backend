package services

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"log/slog"
	"testing"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/guregu/null/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPostgresStorage is a mock implementation of the PostgresStorage interface.
type MockPostgresStorage struct {
	mock.Mock
}

func (m *MockPostgresStorage) GetNewsById(ctx context.Context, id int64) (models.NewsEntity, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(models.NewsEntity), args.Error(1)
}

func (m *MockPostgresStorage) ListNews(ctx context.Context, count int64) ([]models.NewsEntity, error) {
	args := m.Called(ctx, count)
	return (args.Get(0).([]models.NewsEntity)), args.Error(1)
}

func TestGetNewsById(t *testing.T) {
	mockStorage := new(MockPostgresStorage)
	mockLogger := slog.New(slog.NewTextHandler(log.Writer(), nil))
	service := &Service{
		newsStorage: mockStorage,
		log:         mockLogger,
	}

	ctx := context.Background()
	testID := int64(1)

	t.Run("success", func(t *testing.T) {
		mockEntity := models.NewsEntity{
			Id:    testID,
			Title: "Test News",
			Data:  "This is a test news body.",
		}
		mockStorage.On("GetNewsById", ctx, testID).Return(mockEntity, nil)

		expectedResponse := dto.NewsIDItemResponse{
			Id:           testID,
			Title:        "Test News",
			Data:         "This is a test news body.",
			Changed_date: null.Time{NullTime: sql.NullTime{Valid: true}},
		}

		response, err := service.GetNewsById(ctx, testID)

		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, response)
		mockStorage.AssertExpectations(t)
	})

	t.Run("storage error", func(t *testing.T) {
		mockStorage := new(MockPostgresStorage)
		service := &Service{
			newsStorage: mockStorage,
			log:         mockLogger,
		}
		mockStorage.On("GetNewsById", ctx, testID).Return(models.NewsEntity{}, errors.New("storage error"))

		_, err := service.GetNewsById(ctx, testID)

		assert.Error(t, err)
		mockStorage.AssertExpectations(t)
	})

}
