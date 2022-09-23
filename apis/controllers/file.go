package controllers

import (
	"fmt"
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

func (b *FileController) ReadAndInsertDb(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, b.NewResponse(err.Error(), nil))
		return
	}
	fmt.Println(file.Filename)
	open, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, b.NewResponse(err.Error(), nil))
		return
	}
	defer open.Close()
	//reader := csv.NewReader(open)
	//for {
	//	eachRecord, err := reader.Read()
	//	if err != nil || err == io.EOF {
	//		c.JSON(http.StatusInternalServerError, b.NewResponse(err.Error(), nil))
	//		break
	//	}
	//	for value := range eachRecord {
	//		fmt.Printf("%s\n", eachRecord[value])
	//	}
	//}
	//f, err := fileHeader.Open()
	//var p struct {
	//	File multipart.FileHeader `form:"file" binding:"required"`
	//}
	//err := c.MustBindWith(&p, binding.FormMultipart)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, b.NewResponse(err.Error(), nil))
	//	return
	//}
	//file, _ := p.File.Open()

	//for {
	//	each_record, err := file.Read()
	//	if err != nil || err == io.EOF {
	//
	//		log.Fatal(err)
	//		break
	//	}
	//	for value := range each_record {
	//		fmt.Printf("%s\n", each_record[value])
	//	}
	//}
	//fmt.Println(file)
	// file may or may not be in memory
	// I can use file in s3.PutObject as it is a io.Reader but it is from local memory not sreamed from client
	c.JSON(http.StatusOK, b.NewResponse(
		"Success",
		nil,
	))
}
