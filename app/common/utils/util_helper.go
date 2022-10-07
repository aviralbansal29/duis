package commonUtil

import (
	"fmt"
	"strings"

	"github.com/aviralbansal29/duis/log"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func errorMapMessage(validationError validator.FieldError) string {
	switch validationError.Tag() {
	case "required":
		return "This field is required"
	case "numeric":
		return "This field should be numeric"
	case "oneof":
		replacer := *strings.NewReplacer(" ", "/")
		return fmt.Sprintf("Value must be one of (%s)", replacer.Replace(validationError.Param()))
	case "min":
		return fmt.Sprintf("Cannot be less than %s", validationError.Param())
	case "max":
		return fmt.Sprintf("Cannot be more than %s", validationError.Param())
	default:
		log.GetLogger().WithFields(logrus.Fields{"Location": "CommonUtil"}).Error(
			fmt.Sprintf("Unknwon error detected : %s", validationError.Tag()),
		)
		return "Unknown Error"
	}
}
