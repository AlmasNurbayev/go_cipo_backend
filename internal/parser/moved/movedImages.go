package moved

import (
	"log/slog"
	"os"

	cp "github.com/otiai10/copy"
)

func CopyImages(assetsFolder string, newPath string, log *slog.Logger) error {
	op := "moved.CopyImages"
	log = log.With(slog.String("op", op))

	oldPath := newPath + "/import_files"
	newPath = "assets/product_images"
	_, err := os.Stat(oldPath)
	if err != nil {
		log.Error(oldPath + " does not exist")
		return err
	}
	err = cp.Copy(oldPath, newPath)
	if err != nil {
		log.Error(err.Error())
		return err
	} else {
		log.Info(oldPath + " exists and copied successfully to " + newPath)
	}
	return nil
}
