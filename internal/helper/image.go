package helper

import (
	"context"
	"errors"
	"log"
	"mime/multipart"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/sawalreverr/recything/config"
)

func UploadToCloudinary(file interface{}, folderPath string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conf := config.GetConfig().Cloudinary
	cld, err := cloudinary.NewFromParams(conf.CloudName, conf.ApiKey, conf.ApiSecret)
	if err != nil {
		log.Fatal("Cloudinary Enviroment Needed!")
		return "", err
	}

	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: folderPath})
	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}

func ImagesValidation(files []*multipart.FileHeader) ([]multipart.File, error) {
	var response []multipart.File
	for _, file := range files {
		if file.Size > 2*1024*1024 {
			return nil, errors.New("upload image size must less than 2MB")
		}

		fileType := file.Header.Get("Content-Type")
		if !strings.HasPrefix(fileType, "image/") {
			return nil, errors.New("only image allowed")
		}

		src, _ := file.Open()
		defer src.Close()

		response = append(response, src)
	}

	return response, nil
}
