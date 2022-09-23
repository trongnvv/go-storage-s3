package controllers

import (
	"github.com/gin-gonic/gin"
	"go-storage-s3/core/usecases"
	"go-storage-s3/entities"
	"net/http"
)

type FileController struct {
	*baseController
	fileUseCase *usecases.FileUseCase
}

func NewFileController(
	base *baseController,
	fileUseCase *usecases.FileUseCase,
) *FileController {
	return &FileController{
		baseController: base,
		fileUseCase:    fileUseCase,
	}
}

func (b *FileController) GetPresignedUrl(c *gin.Context) {
	url, err := b.fileUseCase.GetPresignedUrl(c, "46082f6c1f7fdf21866e.jpg")
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

func (b *FileController) UploadToS3(c *gin.Context) {
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
