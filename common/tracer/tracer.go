package tracer

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-storage-s3/common/log"
	"go-storage-s3/configs"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"io/ioutil"
)

func InitTracer(cf *configs.Config) *sdktrace.TracerProvider {
	//exp2, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	//if err != nil {
	//	panic(err)
	//}
	tpo := sdktrace.WithResource(resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(cf.ServiceName),
		semconv.ServiceVersionKey.String("v0.1.0"),
		attribute.String("environment", cf.Mode),
	))
	var tp *sdktrace.TracerProvider
	if cf.Jaeger.Active {
		exp1, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(cf.Jaeger.Endpoint)))
		if err != nil {
			panic(err)
		}
		tp = sdktrace.NewTracerProvider( //trace.WithBatcher(exp2),
			sdktrace.WithBatcher(exp1),
			tpo,
		)
	} else {
		tp = sdktrace.NewTracerProvider(
			tpo,
		)
	}

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}))
	return tp
}

func checkIgnorePath(path string) bool {
	listPathIgnore := configs.Get().Jaeger.PathIgnoreLogger
	for _, pathIgnore := range listPathIgnore {
		if pathIgnore == "*" || pathIgnore == path {
			return true
		}
	}
	return false
}

func MidRest(c *gin.Context) {
	//defer func() {
	//	if err := recover(); err != nil {
	//		log.Warnf("mid-rest-trace", err)
	//	}
	//}()

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}
	span := trace.SpanFromContext(c.Request.Context())
	span.SetAttributes(
		attribute.String("body", string(jsonData)),
		attribute.String("query", fmt.Sprintf("%v", c.Request.URL.Query())),
	)
	spanContext := trace.SpanContextFromContext(c.Request.Context())
	path := c.Request.URL.RequestURI()
	if !checkIgnorePath(path) {
		log.Infof("request api", map[string]interface{}{
			"service_name":  configs.Get().ServiceName,
			"uri":           path,
			"request.body":  string(jsonData),
			"request.query": c.Request.URL.Query(),
			"trace_id":      spanContext.TraceID(),
		})
	}

}

func MidGrpcClient() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, resp interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		//defer func() {
		//	if err := recover(); err != nil {
		//		log.Warnf("mid-grpc-client-trace", err)
		//	}
		//}()

		span := trace.SpanFromContext(ctx)
		span.SetAttributes(attribute.String("request", fmt.Sprintf("%v", req)))
		spanContext := trace.SpanContextFromContext(ctx)
		path := method
		if !checkIgnorePath(path) {
			log.Infof("grpc client request", map[string]interface{}{
				"service_name": configs.Get().ServiceName,
				"method":       path,
				"request":      req,
				"trace_id":     spanContext.TraceID(),
			})
		}

		return invoker(ctx, method, req, resp, cc, opts...)
	}
}

func MidGrpcServer() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		//defer func() {
		//	if err := recover(); err != nil {
		//		log.Warnf("mid-grpc-server-trace", err)
		//	}
		//}()
		span := trace.SpanFromContext(ctx)
		span.SetAttributes(
			attribute.String("request", fmt.Sprintf("%v", req)),
			attribute.String("response", fmt.Sprintf("%v", resp)),
		)

		spanContext := trace.SpanContextFromContext(ctx)
		path := info.FullMethod
		if !checkIgnorePath(path) {
			log.Infof("grpc server resp", map[string]interface{}{
				"service_name": configs.Get().ServiceName,
				"method":       info.FullMethod,
				"request":      req,
				"response":     resp,
				"trace_id":     spanContext.TraceID(),
			})
		}
		return resp, err
	}
}
