package fxloader

import (
	"go-storage-s3/adapters/pgsql"
	dbRepo "go-storage-s3/adapters/pgsql/repositories"
	"go-storage-s3/adapters/s3"
	s3Repo "go-storage-s3/adapters/s3/repositories"
	"go-storage-s3/apis/controllers"
	"go-storage-s3/apis/middlewares"
	"go-storage-s3/apis/routers"
	"go-storage-s3/core/usecases"
	"go.uber.org/fx"
)

func Load() []fx.Option {
	return []fx.Option{
		fx.Options(loadApis()...),
		fx.Options(loadAdapter()...),
		fx.Options(loadUseCase()...),
	}
}

func loadApis() []fx.Option {
	return []fx.Option{
		fx.Provide(routers.NewApiRouter),
		fx.Provide(middlewares.NewMiddleWare),
		fx.Provide(controllers.NewBaseController),
		fx.Provide(controllers.NewFileController),
	}
}

func loadUseCase() []fx.Option {
	return []fx.Option{
		fx.Provide(usecases.NewFileS3UseCase),
		fx.Provide(usecases.NewReadFileCSVHandleUseCase),
	}
}

func loadAdapter() []fx.Option {
	return []fx.Option{
		fx.Provide(s3.Connect),
		fx.Provide(pgsql.Connect),
		fx.Provide(dbRepo.NewFileRepository),
		fx.Provide(dbRepo.NewFileCsvRepository),
		fx.Provide(s3Repo.NewFileRepository),
	}
}
