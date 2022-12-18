package middlewares

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-product-service/config"
	"github.com/thirumathikart/thirumathikart-product-service/generated/user"
	"github.com/thirumathikart/thirumathikart-product-service/rpcs"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func Authenticator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		headers := c.Request().Header
		userToken := headers.Get("Authorization")
		log.Println(userToken)
		if userToken == "" {
			return SendResponse(c, http.StatusServiceUnavailable, "Unable to authorize try later")
		}
		md := metadata.New(map[string]string{"secret": "xxxx"})
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		var opts []grpc.DialOption
		opts = append(
			opts,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			WithClientUnaryInterceptor())
		conn, err := grpc.Dial(config.AuthService, opts...)
		if err != nil {
			log.Println(err)
			return SendResponse(c, http.StatusServiceUnavailable, "Unable to authorize try later")
		}
		defer conn.Close()
		client := user.NewUserServiceClient(conn)
		response, err := rpcs.AuthRPC(ctx, userToken, client)
		if err != nil {
			return SendResponse(c, http.StatusBadRequest, "Error Occurred")
		}
		if !response.IsSuccess {
			return SendResponse(c, http.StatusUnauthorized, "Unauthorized")
		}
		c.Set("user", response.User)
		return next(c)
	}
}
