package usecases

import (
	"context"
	"go-storage-s3/core/ports"
)

type FileS3UseCase struct {
	fileDatabaseRepository ports.FileDatabaseRepository
	fileS3RepositoryPort   ports.FileS3Repository
}

func NewFileS3UseCase(
	fileDatabaseRepository ports.FileDatabaseRepository,
	fileS3RepositoryPort ports.FileS3Repository,
) *FileS3UseCase {
	return &FileS3UseCase{
		fileDatabaseRepository: fileDatabaseRepository,
		fileS3RepositoryPort:   fileS3RepositoryPort,
	}
}

func (u *FileS3UseCase) Upload() {
	u.fileS3RepositoryPort.Upload()
}
func (u *FileS3UseCase) GetPresignedUrl(ctx context.Context, path string) (string, error) {
	res, err := u.fileS3RepositoryPort.GetPresignedUrl(ctx, path)
	if err != nil {
		return "", err
	}
	return res, nil
	//u.fileDatabaseRepository.Save()
	//return ""
}
