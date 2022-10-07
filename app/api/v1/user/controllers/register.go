package controllers

import (
	"fmt"

	"github.com/aviralbansal29/duis/app/api/v1/user/serializers"
	"github.com/aviralbansal29/duis/app/api/v1/user/services"
	"github.com/labstack/echo/v4"
)

// Register godoc
// @Summary       Register User
// @Description   Registers the user with default variant
// @Tags          User
// @Security      UserID
// @Accept        json
// @Produce       json
// @Param         data  body  serializers.RegisterUserRequestBody true  "Body"
// @Success       200  {object}  serializers.RegisterUserResponse
// @Failure       400
// @Failure       404
// @Failure       500
// @Router        /api/v1/users [post]
func Register(c echo.Context) error {
	fmt.Println("Registering user")
	data, err := serializers.SerializeRegisterUser(c)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(data)

	responseData, err := services.RegisterUser(data)
	if err != nil {
		return err
	}

	return serializers.DeserializeRegisterUser(c, responseData)
}
