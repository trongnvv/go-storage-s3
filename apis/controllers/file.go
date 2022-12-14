package controllers

import (
	"github.com/gin-gonic/gin"
	"go-storage-s3/core/usecases"
	"go-storage-s3/entities"
	"mime/multipart"
	"net/http"
)

type FileController struct {
	*baseController
	fileS3UseCase            *usecases.FileS3UseCase
	readFileCSVHandleUseCase *usecases.ReadFileCSVHandleUseCase
}

func NewFileController(
	base *baseController,
	fileS3UseCase *usecases.FileS3UseCase,
	readFileCSVHandleUseCase *usecases.ReadFileCSVHandleUseCase,
) *FileController {
	return &FileController{
		baseController:           base,
		fileS3UseCase:            fileS3UseCase,
		readFileCSVHandleUseCase: readFileCSVHandleUseCase,
	}
}

func (b *FileController) S3Upload(c *gin.Context) {
	b.fileS3UseCase.Upload()
	c.JSON(http.StatusOK, b.NewResponse(
		"Success",
		nil,
	))
}
func (b *FileController) GetPresignedUrl(c *gin.Context) {
	url, err := b.fileS3UseCase.GetPresignedUrl(c, "46082f6c1f7fdf21866e.jpg")
	if err != nil {
		c.JSON(http.StatusInternalServerError, b.NewResponse(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, b.NewResponse(
		"Success",
		entities.ResponseFile{
			Name: url,
		},
	))
}

func (b *FileController) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, b.NewResponse(err.Error(), entities.ResponseFile{Name: file.Filename}))
		return
	}
	err = c.SaveUploadedFile(file, "upload/"+file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, b.NewResponse(err.Error(), entities.ResponseFile{Name: file.Filename}))
		return
	}

	c.JSON(http.StatusOK, b.NewResponse(
		"Success",
		entities.ResponseFile{
			Name: file.Filename,
			Size: file.Size,
		},
	))
}

func (b *FileController) ReadAndInsertDb(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, b.NewResponse(err.Error(), nil))
		return
	}

	go func(file *multipart.FileHeader, readFileCSVHandleUseCase *usecases.ReadFileCSVHandleUseCase) {
		open, err := file.Open()
		defer open.Close()
		if err != nil {
			return
		}
		readFileCSVHandleUseCase.Run(open)
	}(file, b.readFileCSVHandleUseCase)

	c.JSON(http.StatusOK, b.NewResponse(
		"Success",
		entities.ResponseFile{
			Name: file.Filename,
			Size: file.Size,
		},
	))
}
