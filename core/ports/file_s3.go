package ports

import (
	"context"
)

type FileS3Repository interface {
	GetPresignedUrl(ctx context.Context, path string) (string, error)
}
