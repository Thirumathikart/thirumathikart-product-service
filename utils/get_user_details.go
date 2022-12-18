package utils

import (
	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-product-service/generated/user"
)

func GetUserDetails(c echo.Context) (*user.User, error) {
	userDetails := c.Get("user").(*user.User)
	return userDetails, nil
}
