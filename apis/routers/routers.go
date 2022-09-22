package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"go-storage-s3/apis/controllers"
	"go-storage-s3/apis/middlewares"
	"go-storage-s3/common/tracer"
	"go-storage-s3/configs"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	_ "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type ApiRouter struct {
	Engine *gin.Engine
}

func NewApiRouter(
	logger *zerolog.Logger,
	cf *configs.Config,
	mid *middlewares.MiddleWare,
	fileCtrl *controllers.FileController,
) *ApiRouter {
	engine := gin.New()
	gin.DisableConsoleColor()
	engine.Use(gin.Logger())
	engine.Use(otelgin.Middleware(cf.ServiceName))
	engine.Use(tracer.MidRest)
	engine.Use(mid.Recovery(logger))
	r := engine.RouterGroup.Group(cf.Prefix)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.Use(mid.Authenticate)
	{
		r.GET("/presigned_url", fileCtrl.GetPresignedUrl)

		r.POST("/upload", fileCtrl.Upload)
		r.POST("/upload-to-s3", fileCtrl.UploadToS3)
	}
	return &ApiRouter{
		Engine: engine,
	}
}
