package usecases

import (
	"encoding/csv"
	"go-storage-s3/core/ports"
	"go-storage-s3/entities/models"
	"io"
)

type ReadFileCSVHandleUseCase struct {
	fileCsvDatabaseRepository ports.FileCsvDatabaseRepository
	fileDatabaseRepository    ports.FileDatabaseRepository
	fileS3RepositoryPort      ports.FileS3Repository
}

func NewReadFileCSVHandleUseCase(
	fileCsvDatabaseRepository ports.FileCsvDatabaseRepository,
	fileDatabaseRepository ports.FileDatabaseRepository,
	fileS3RepositoryPort ports.FileS3Repository,
) *ReadFileCSVHandleUseCase {
	return &ReadFileCSVHandleUseCase{
		fileCsvDatabaseRepository: fileCsvDatabaseRepository,
		fileDatabaseRepository:    fileDatabaseRepository,
		fileS3RepositoryPort:      fileS3RepositoryPort,
	}
}

func (u *ReadFileCSVHandleUseCase) Run(open io.Reader) {
	reader := csv.NewReader(open)
	index := 0
	for {
		line, err := reader.Read()
		if err != nil || err == io.EOF {
			return
		}
		if index == 0 {
			index++
			continue
		}
		u.insert(line)
		//for _, col := range lines {
		//	fmt.Println(col)
		//}
	}
}

func (u *ReadFileCSVHandleUseCase) insert(line []string) {
	// validate
	data := &models.DataCsvModel{
		Description: line[0],
		Industry:    line[1],
		Level:       line[2],
		Size:        line[3],
		LineCode:    line[4],
		Value:       line[5],
	}
	err := u.fileCsvDatabaseRepository.Save(data)
	if err != nil {
		return
	}
}
