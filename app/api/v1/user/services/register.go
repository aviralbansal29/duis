package services

import (
	"net/http"

	"github.com/aviralbansal29/duis/app/api/v1/user/serializers"
	"github.com/aviralbansal29/duis/app/models"
	"github.com/aviralbansal29/duis/config"
	"github.com/aviralbansal29/duis/constant"
	"github.com/aviralbansal29/duis/log"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// RegisterUser registers the user
func RegisterUser(data serializers.RegisterUserRequestBody) (models.User, error) {
	user := models.User{Name: data.Name, Email: data.Email, VariantID: constant.DefaultVariantID}
	var existingCount int64
	err := config.DatabaseHandler().Model(&models.User{}).Where("email = ?", data.Email).Count(&existingCount).Error
	if err != nil {
		log.GetLogger().WithFields(logrus.Fields{"Location": "Unique Email Check"}).Error(err.Error())
		return user, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	if existingCount != 0 {
		return user, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"errors": map[string]string{"email": "Account already exists with this email"},
		})
	}
	err = config.DatabaseHandler().Create(&user).Error
	if err != nil {
		log.GetLogger().WithFields(logrus.Fields{"Location": "Creating User"}).Error(err.Error())
		return user, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return user, nil
}
