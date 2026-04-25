package parserJSON

import (
	"archive/zip"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/config"
)

// находим в папке Input ZIP-архив, перемещаем в папку input json и картинки
func MovedFTPFiles(cfg *config.Config, log *slog.Logger) (string, error) {
	op := "parserJSON.MovedInputFiles"
	log = log.With(slog.String("op", op))

	ftpfolder := cfg.ParserJSON.PARSER_FTP_PATH
	inputPath := cfg.Parser.PARSER_INPUT_PATH
	prefix := cfg.ParserJSON.PARSER_FILE_PREFIX

	// 1. Ищем последний ZIP файл
	latestZip, err := findLatestZip(ftpfolder, prefix)
	if err != nil {
		log.Error("Error find latest zip file:", slog.String("error", err.Error()))
		return "", err
	}
	if latestZip == "" {
		log.Error("Error zip file not found")
		return "", fmt.Errorf("zip file not found")
	}

	log.Info("Found latest zip file: ", slog.String("latestZip", latestZip))

	// 2. Распаковываем
	jsonPath, err := unzipAndFindJSON(latestZip, inputPath)
	if err != nil {
		log.Error("Error unzip:", slog.String("error", err.Error()))
		return "", err
	}
	log.Info("Unzip successfully")

	// TODO отключено
	// 3. Удаляем исходный архив после успешной распаковки
	//	err = os.Remove(latestZip)
	//	if err != nil {
	//		log.Warn("Error removing zip file:", slog.String("error", err.Error()))
	//	}
	//	log.Info("Zip file removed successfully: ", slog.String("zipPath", latestZip))

	return jsonPath, nil
}

// findLatestZip ищет самый свежий файл по дате изменения
func findLatestZip(dir, prefix string) (string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}

	var zipFiles []os.FileInfo
	for _, entry := range entries {
		// Проверяем префикс, расширение и что это не папка
		if !entry.IsDir() && strings.HasPrefix(entry.Name(), prefix) && strings.HasSuffix(strings.ToLower(entry.Name()), ".zip") {
			info, err := entry.Info()
			if err == nil {
				zipFiles = append(zipFiles, info)
			}
		}
	}

	if len(zipFiles) == 0 {
		return "", nil
	}

	// Сортируем по времени изменения (от новых к старым)
	sort.Slice(zipFiles, func(i, j int) bool {
		return zipFiles[i].ModTime().After(zipFiles[j].ModTime())
	})

	return filepath.Join(dir, zipFiles[0].Name()), nil
}

func unzipAndFindJSON(src, dest string) (string, error) {
	r, err := zip.OpenReader(src)
	if err != nil {
		return "", err
	}
	defer r.Close()

	var foundJSONPath string
	jsonCount := 0

	// Сначала создаем целевую директорию (абсолютный путь)
	absDest, err := filepath.Abs(dest)
	if err != nil {
		return "", err
	}

	for _, f := range r.File {
		fpath := filepath.Join(absDest, f.Name)

		// Проверка ZipSlip
		if !strings.HasPrefix(fpath, filepath.Clean(absDest)+string(os.PathSeparator)) {
			return "", fmt.Errorf("некорректный путь в архиве: %s", fpath)
		}

		// Если это файл JSON, запоминаем его
		if !f.FileInfo().IsDir() && strings.HasSuffix(strings.ToLower(f.Name), ".json") {
			foundJSONPath = fpath
			jsonCount++
		}

		// Распаковка
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return "", err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return "", err
		}

		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return "", err
		}

		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()

		if err != nil {
			return "", err
		}
	}

	// Валидация количества JSON файлов
	if jsonCount == 0 {
		return "", fmt.Errorf("в архиве не найдено ни одного JSON файла")
	}
	if jsonCount > 1 {
		return "", fmt.Errorf("в архиве найдено несколько JSON файлов (%d), ожидался один", jsonCount)
	}

	return foundJSONPath, nil
}
