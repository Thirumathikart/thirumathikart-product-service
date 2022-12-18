package middlewares

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc/grpclog"
)

func InitLogger(server *echo.Echo) {
	server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} | ${method} ${uri} \t | ${latency_human}\n",
		Output: server.Logger.Output(),
	}))
}

var GrpcLog grpclog.LoggerV2

func GrpcLogger() {
	GrpcLog = grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr)
	grpclog.SetLoggerV2(GrpcLog)
}
