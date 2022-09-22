package controllers

import (
	"github.com/go-playground/validator/v10"
	"go-storage-s3/common/log"
)

type ResponseResource struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

type baseController struct {
	validate *validator.Validate
}

func NewBaseController() *baseController {
	return &baseController{
		validate: validator.New(),
	}
}

func (b *baseController) NewResponse(message string, body interface{}) *ResponseResource {
	return &ResponseResource{
		Message: message,
		Body:    body,
	}
}

func (b *baseController) validateRequest(request interface{}) error {
	err := b.validate.Struct(request)
	if err != nil {
		for _, errValidate := range err.(validator.ValidationErrors) {
			log.Debugf("query invalid, err:[%s]", errValidate)
		}
		return err
	}
	return nil
}
