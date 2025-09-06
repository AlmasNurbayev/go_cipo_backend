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

func (s *Storage) CreateVidModeli(ctx context.Context, data models.VidModeliEntity) (int64, error) {
	op := "postgres.CreateVidModeli"
	log := s.log.With(slog.String("op", op))

	query := `INSERT INTO vid_modeli
	(id_1c, name_1c, registrator_id) 
		VALUES 
		($1, $2, $3) 
		RETURNING id;`
	db := *s.Tx

	err := db.QueryRow(ctx, query,
		data.Id_1c,
		data.Name_1c,
		data.Registrator_id).Scan(
		&data.Id)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return data.Id, err
	}
	return data.Id, nil
}

func (s *Storage) ListVidModeli(ctx context.Context) ([]models.VidModeliEntity, error) {
	op := "postgres.ListVidModeli"
	log := s.log.With("op", op)

	var vidsModeli = []models.VidModeliEntity{}
	query := `SELECT id, id_1c, name_1c, registrator_id FROM vid_modeli;`
	var err error
	// если есть транзакция, используем ее, иначе стандартный пул
	if s.Tx != nil {
		db := *s.Tx
		err = pgxscan.Select(ctx, db, &vidsModeli, query)
	} else {
		db := s.Db
		err = pgxscan.Select(ctx, db, &vidsModeli, query)
	}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// если выкидывается ошибка нет строк, возвращаем пустой массив
			return vidsModeli, nil
		}
		log.Error(err.Error())
		return vidsModeli, errorsShare.ErrInternalError.Error
	}
	return vidsModeli, nil
}

func (s *Storage) UpdateVidModeli(ctx context.Context, data models.VidModeliEntity) error {
	op := "postgres.UpdateVidModeli"
	log := s.log.With(slog.String("op", op))

	query := `UPDATE vid_modeli SET
	id_1c = $1, name_1c = $2, registrator_id = $3 
		WHERE id_1c = $4 RETURNING *;`
	db := *s.Tx

	err := pgxscan.Get(ctx, db, &data, query,
		data.Id_1c,
		data.Name_1c,
		data.Registrator_id, data.Id_1c)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return err
	}
	return nil
}

func (s *Storage) GetVidModeliById(ctx context.Context, id int64) (models.VidModeliEntity, error) {
	op := "postgres.GetVidModeliById"
	log := s.log.With(slog.String("op", op))

	query := `SELECT * FROM vid_modeli WHERE id = $1;`
	db := s.Db

	data := models.VidModeliEntity{}

	err := pgxscan.Get(ctx, db, &data, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return data, errorsShare.ErrInternalError.Error
		}
		log.Error("error: ", slog.String("err", err.Error()))
		return data, err
	}
	return data, nil
}

func (s *Storage) ListVidModeliIdExcludeNames(ctx context.Context, exclude []string) ([]int64, error) {
	op := "postgres.ListVidModeli"
	log := s.log.With("op", op)

	var vidsModeli = []models.VidModeliEntity{}
	var ids = []int64{}
	query := `SELECT id FROM vid_modeli
	WHERE name_1c = ANY($1)
	;`
	var err error
	// если есть транзакция, используем ее, иначе стандартный пул
	if s.Tx != nil {
		db := *s.Tx
		err = pgxscan.Select(ctx, db, &vidsModeli, query, exclude)
	} else {
		db := s.Db
		err = pgxscan.Select(ctx, db, &vidsModeli, query, exclude)
	}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// если выкидывается ошибка нет строк, возвращаем пустой массив
			return ids, nil
		}
		log.Error(err.Error())
		return ids, errorsShare.ErrInternalError.Error
	}
	for _, r := range vidsModeli {
		ids = append(ids, r.Id)
	}

	return ids, nil
}
