package moved

import (
	"log/slog"
	"os"
	"time"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/config"
)

// находим в папке Input файлы, создаем папку Webdata_дата_время,
// перемещаем туда offers0_1.xml и import0_1.xml
func MovedInputFiles(cfg *config.Config, log *slog.Logger) (*MovedInputFilesT, error) {
	op := "moved.MovedInputFiles"
	log = log.With(slog.String("op", op))

	currentTime := time.Now()
	folderName := "input/webdata_" + currentTime.Format("2006_01_02_15_04_05")
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		err := os.Mkdir(folderName, 0755)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
	}

	filesName := []InputFilesT{
		{"classificator", cfg.Parser.PARSER_CLASSIFICATOR_FILE},
		{"offer", cfg.Parser.PARSER_OFFER_FILE},
		{"imageFolder", cfg.Parser.PARSER_IMAGE_FOLDER},
	}

	for i, file := range filesName {
		oldPath := "input/" + file.PathFile
		newPath := folderName + "/" + file.PathFile
		filesName[i].PathFile = newPath
		if _, err := os.Stat(oldPath); err == nil {
			err := os.Rename(oldPath, newPath)
			if err != nil {
				if i == 2 {
					// если это imageFolder, то не прерываем программу
					log.Error("Error image folder moved:", slog.String("error", err.Error()))
				} else {
					log.Error("Error moving file:", slog.String("error", err.Error()))
					return nil, err
				}
			} else {
				log.Info(file.PathFile + " exists and moved successfully")
			}
		} else {
			if i == 2 {
				// если это imageFolder, то не прерываем программу
				log.Error(file.PathFile + " does not exist")
			} else {
				log.Error(file.PathFile + " does not exist")
				return nil, err
			}
		}
	}
	return &MovedInputFilesT{Files: filesName, NewPath: folderName}, nil
}

type InputFilesT struct {
	TypeFile string
	PathFile string
}

type MovedInputFilesT struct {
	Files   []InputFilesT
	NewPath string
}
