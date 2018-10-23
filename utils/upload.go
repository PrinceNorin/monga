package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/PrinceNorin/monga/config"
	"github.com/gin-gonic/gin"
)

func UploadFile(file *multipart.FileHeader, c *gin.Context) (string, error) {
	ext := filepath.Ext(file.Filename)
	path := filepath.Join(
		config.Get().Upload.Dir,
		fmt.Sprintf("%s%s", RandString(16), ext),
	)
	src, err := file.Open()
	if err != nil {
		return "", nil
	}
	defer src.Close()

	dest, _ := os.Create(path)
	if err != nil {
		return "", nil
	}
	defer dest.Close()
	io.Copy(dest, src)

	return path, nil
}
