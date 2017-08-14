package pop

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/satori/go.uuid"
)

var CachedModel map[string]interface{}
var CacheEnabled = true

func ModelChanges(model interface{}) (map[string]interface{}, error) {
	cached := reflect.New(reflect.ValueOf(model))

}

func GetCachedModel(model interface{}, id interface{}) error {
	val := reflect.ValueOf(id)
	var idStr string
	if val.Type().Name() == "uuid.UUID" {
		idStr = val.Interface().(uuid.UUID).String()
	} else {
		idStr = string(val.Interface().(int64))
	}
	name := reflect.TypeOf(model).Elem().Name()
	key := name + "-" + idStr
	res, ok := CachedModel[key]
	if !ok {
		return fmt.Errorf("model missing: %v", key)
	}
	model = &res
	return nil
}

func CacheModel(model interface{}) error {
	name := reflect.TypeOf(model).Elem().Name()
	if reflect.ValueOf(model).Kind() == reflect.Ptr {
		return fmt.Errorf("CacheModel must be a value, not pointer")
	}
	// models, ok := CachedModel[name]

	field := reflect.ValueOf(model).Elem().FieldByName("ID")
	var id string
	if field.Type().Name() == "uuid.UUID" {
		id = field.Interface().(uuid.UUID).String()
	} else {
		id = string(field.Interface().(int64))
	}

	CachedModel[name+"-"+id] = model
	return nil
}
