package postgres

import (
	"context"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (s *Storage) CreateRegistrator(ctx context.Context, data models.RegistratorEntity) (int64, error) {
	op := "postgres.CreateRegistrator"
	log := s.log.With(slog.String("op", op))

	query := `INSERT INTO registrator
	(operation_date, name_folder, name_file, user_id, date_schema,
		id_catalog, id_class, name_catalog, name_class, ver_schema, only_change) 
		VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) 
		RETURNING id;`
	db := *s.Tx

	err := db.QueryRow(ctx, query,
		data.Operation_date,
		data.Name_folder, data.Name_file,
		data.User_id, data.Date_schema,
		data.Id_catalog, data.Id_class,
		data.Name_catalog, data.Name_class, data.Ver_schema,
		data.Only_change).Scan(
		&data.Id)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return data.Id, err
	}
	return data.Id, nil
}

func (s *Storage) GetRegistratorById(ctx context.Context, id int64) (models.RegistratorEntity, error) {
	op := "postgres.GetRegistratorById"
	log := s.log.With(slog.String("op", op))

	query := `SELECT id, operation_date, name_folder, name_file, user_id, date_schema,
		id_catalog, id_class, name_catalog, name_class, ver_schema, only_change 
		FROM registrator WHERE id = $1;`
	db := *s.Tx

	var registrator models.RegistratorEntity
	err := pgxscan.Get(ctx, db, &registrator, query, id)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return registrator, err
	}
	return registrator, nil
}

func (s *Storage) GetLastOfferRegistrator(ctx context.Context) (models.RegistratorEntity, error) {
	op := "postgres.GetRegistratorById"
	log := s.log.With(slog.String("op", op))

	query := `SELECT id, operation_date, name_folder, name_file, user_id, date_schema,
		id_catalog, id_class, name_catalog, name_class, ver_schema, only_change 
		FROM registrator WHERE name_file LIKE '%offer%' ORDER BY id desc LIMIT 1`
	db := s.Db

	var registrator models.RegistratorEntity
	err := pgxscan.Get(ctx, db, &registrator, query)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return registrator, err
	}
	return registrator, nil
}
