package commonUtil

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

// GetStringMapStringList returns []map[string]string from []interface{}
func GetStringMapStringList(dataInterface []interface{}) []map[string]string {
	var dataMap []map[string]string
	var strKey, strVal string
	dataMapVal := make(map[string]string)
	for _, data := range dataInterface {
		for key, val := range data.(map[interface{}]interface{}) {
			strKey, strVal = fmt.Sprintf("%v", key), fmt.Sprintf("%v", val)
			dataMapVal[strKey] = strVal
		}
		dataMap = append(dataMap, dataMapVal)
		dataMapVal = map[string]string{}
	}
	return dataMap
}

// GetErrorMap returns errors in a map[string]string
func GetErrorMap(paramType reflect.Type, err error, mapType ...string) map[string]string {
	respErr := make(map[string]string)
	if len(mapType) == 0 {
		mapType = []string{"query"}
	}
	switch err.(type) {
	case validator.ValidationErrors:
		for _, validationError := range err.(validator.ValidationErrors) {
			field, _ := paramType.Elem().FieldByName(validationError.Field())
			fieldName, _ := field.Tag.Lookup(mapType[0])
			respErr[fieldName] = errorMapMessage(validationError)
		}
	default:
		respErr["Unknown field error"] = err.Error()
	}
	return respErr
}

// ConvertType Converts one data type to another using JSON
func ConvertType(src interface{}, des interface{}) error {
	tempJSON, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tempJSON, des)
	return err
}
