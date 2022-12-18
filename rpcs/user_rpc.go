package rpcs

import (
	"context"

	"github.com/thirumathikart/thirumathikart-product-service/generated/user"
)

func AuthRPC(userToken string, client user.UserServiceClient) (*user.AuthResponse, error) {

	return client.AuthRPC(context.Background(),
		&user.AuthRequest{
			UserToken: userToken,
		})
}

func UserRPC(userID uint, client user.UserServiceClient) (*user.UserResponse, error) {

	return client.UserRPC(context.Background(),
		&user.UserRequest{
			UserID: uint32(userID),
		})
}
