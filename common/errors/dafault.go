package errors

import (
	"gorm.io/gorm"
)

var (
	ErrUnAuthorized        = NewCustomHttpError(Unauthorized, "unauthorized")
	ErrSignedData          = NewCustomHttpError(Unauthorized, "check_signed_data")
	ErrMissingMerchantId   = NewCustomHttpError(BadRequest, "missing_merchant_id")
	ErrMissingMerchantCode = NewCustomHttpError(BadRequest, "missing_merchant_code")
	ErrTokenExpires        = NewCustomHttpError(Unauthorized, "token_expires")
	ErrPasswordWrong       = NewCustomHttpError(PasswordWrong, "password_wrong")
)

var (
	ErrSystem           = NewCustomHttpError(SystemError, "system_error")
	ErrBadRequest       = NewCustomHttpError(BadRequest, "invalid_request")
	ErrResourceNotFound = NewCustomHttpError(ResourceNotFound, "resource_not_found")
	ErrConflict         = NewCustomHttpError(Conflict, "username_has_exist ")
)

var (
	ErrEntityNotFound = gorm.ErrRecordNotFound
)

var (
	ErrWebhookNotSupportType = NewCustomHttpError(BadRequest, "not_support_webhook_type")
	ErrWebhookWrongPayload   = NewCustomHttpError(BadRequest, "wrong_payload_webhook")
)

var (
	ErrMerchantNotExist = NewCustomHttpError(ResourceNotFound, "merchant_not_exist")
	ErrMerchantExpired  = NewCustomHttpError(MerchantExpired, "merchant_expired")
)
