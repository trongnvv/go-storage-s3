package controllers

import (
	"github.com/gin-gonic/gin"
	"go-storage-s3/entities"
	"net/http"
)

type FileController struct {
	*baseController
}

func NewFileController(
	base *baseController,
) *FileController {
	return &FileController{
		baseController: base,
	}
}

func (b *FileController) GetPresignedUrl(c *gin.Context) {
	//var getListBankReq requests.GetListBankRequest
	//if err := c.ShouldBindQuery(&getListBankReq); err != nil {
	//	log.Warnf("bind query err, err:[%s]", err)
	//	b.DefaultBadRequest(c)
	//	return
	//}
	//if err := b.validateRequest(getListBankReq); err != nil {
	//	b.BadRequest(c, err.Error())
	//	return
	//}
	//data, customErr := b.getListBankUseCase.GetBankInfos(c.Request.Context())
	//if customErr != nil {
	//	b.ErrorData(c, customErr)
	//	return
	//}

	c.JSON(http.StatusOK, nil)
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
