package fxloader

import (
	"go-storage-s3/adapters/pgsql"
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
		fx.Provide(usecases.NewFileUseCase),
	}
}

func loadAdapter() []fx.Option {
	return []fx.Option{
		fx.Provide(pgsql.DatabaseConnect),
	}
}
