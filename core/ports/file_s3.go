package ports

import (
	"context"
)

type FileS3Repository interface {
	Upload()
	GetPresignedUrl(ctx context.Context, path string) (string, error)
}
