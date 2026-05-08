package services

import (
	"context"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/errorsShare"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/lib/utils"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/guregu/null/v5"
	"github.com/jinzhu/copier"
)

func (s *Service) GetUserByIdService(ctx context.Context, id int64) (dto.UserResponse, error) {
	op := "services.GetUserByIdService"
	log := s.log.With(slog.String("op", op))

	userDTO := dto.UserResponse{}

	userEntity, err := s.userStorage.GetUserByIdStorage(ctx, id)
	if err != nil {
		log.Warn(err.Error())
		if err == errorsShare.ErrUserNotFound.Error {
			return userDTO, errorsShare.ErrUserNotFound.Error
		}
		return userDTO, err
	}

	err = copier.Copy(&userDTO, &userEntity)
	if err != nil {
		log.Warn(err.Error())
		return userDTO, errorsShare.ErrInternalError.Error
	}

	return userDTO, nil
}

func (s *Service) GetUserByNameService(ctx context.Context, name string) (dto.UserResponse, error) {
	op := "services.GetUserByNameService"
	log := s.log.With(slog.String("op", op))

	userEntity, err := s.userStorage.GetUserByNameStorage(ctx, name)
	userDTO := dto.UserResponse{}

	if err != nil {
		log.Warn(err.Error())
		if err == errorsShare.ErrUserNotFound.Error {
			return userDTO, errorsShare.ErrUserNotFound.Error
		}
		return userDTO, err
	}

	err = copier.Copy(&userDTO, &userEntity)
	if err != nil {
		log.Warn(err.Error())
		return userDTO, err
	}
	return userDTO, nil
}

func (s *Service) CreateUserService(ctx context.Context, user dto.AuthRegisterRequest) (dto.AuthRegisterResponse, error) {
	op := "services.CreateUserService"
	log := s.log.With(slog.String("op", op))

	response := dto.AuthRegisterResponse{}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Error("error hash password", slog.String("err", err.Error()))
		return response, err
	}
	user.Password = hashedPassword

	userEntity := models.UserEntity{}
	err = copier.Copy(&userEntity, &user)
	if err != nil {
		log.Warn(err.Error())
		return response, errorsShare.ErrInternalError.Error
	}
	userEntity.Role = null.StringFrom("user")

	id, err := s.userStorage.CreateUserStorage(ctx, userEntity)
	if err != nil {
		log.Warn(err.Error())
		if err == errorsShare.ErrUserNotFound.Error {
			return response, errorsShare.ErrUserNotFound.Error
		}
		return response, err
	}

	response.Id = id
	response.Email = user.Email

	return response, nil
}
