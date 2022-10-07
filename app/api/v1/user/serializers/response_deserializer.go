package serializers

import (
	"net/http"

	"github.com/aviralbansal29/duis/app/models"
	"github.com/labstack/echo/v4"
)

type RegisterUserResponse struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	VariantName string `json:"variant"`
}

func DeserializeRegisterUser(ctx echo.Context, data models.User) error {
	return ctx.JSON(http.StatusCreated, RegisterUserResponse{
		Name: data.Name, Email: data.Email, VariantName: data.Variant.Name,
	})
}
