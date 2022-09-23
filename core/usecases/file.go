package usecases

import (
	"context"
	"go-storage-s3/core/ports"
)

type FileUseCase struct {
	fileDatabaseRepository ports.FileDatabaseRepository
	fileS3RepositoryPort   ports.FileS3Repository
}

func NewFileUseCase(
	fileDatabaseRepository ports.FileDatabaseRepository,
	fileS3RepositoryPort ports.FileS3Repository,
) *FileUseCase {
	return &FileUseCase{
		fileDatabaseRepository: fileDatabaseRepository,
		fileS3RepositoryPort:   fileS3RepositoryPort,
	}
}

func (u *FileUseCase) GetPresignedUrl(ctx context.Context, path string) (string, error) {
	res, err := u.fileS3RepositoryPort.GetPresignedUrl(ctx, path)
	if err != nil {
		return "", err
	}
	return res, nil
	//u.fileDatabaseRepository.Save()
	//return ""
}
