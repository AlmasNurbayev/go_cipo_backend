package services

import (
	"context"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/errorsShare"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/lib/utils"
)

func (s *Service) LoginUserService(ctx context.Context, user dto.AuthLoginRequest) (dto.AuthLoginResponse, error) {
	op := "services.LoginUserService"
	log := s.log.With(slog.String("op", op))

	response := dto.AuthLoginResponse{}

	userEntity, err := s.userStorage.GetUserByEmailStorage(ctx, user.Email)
	if err != nil {
		log.Warn(err.Error())
		if err == errorsShare.ErrUsernameOrPasswordIsWrong.Error {
			return response, errorsShare.ErrUsernameOrPasswordIsWrong.Error
		}
		return response, err
	}
	if !userEntity.Password.Valid {
		log.Warn("empty password", "user_id", userEntity.Id)
		return response, errorsShare.ErrUsernameOrPasswordIsWrong.Error
	}
	if utils.CheckPassword(userEntity.Password.String, user.Password) != nil {
		log.Warn("password error", "user_id", userEntity.Id)
		return response, errorsShare.ErrUsernameOrPasswordIsWrong.Error
	}
	response.Id = userEntity.Id

	return response, nil
}
