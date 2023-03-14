package steps

import (
	"reflect"
	"strings"

	"github.com/bitrise-io/envman/models"
)

type XPut struct {
	Identifier string
	Options    XPutOptions
	Value      any
}

type XPutOptions struct {
	Description string
	Summary     string
	Title       string

	IsRequired bool
}

func createYamlTagToFieldMap(t reflect.Type) map[string]reflect.StructField {
	total := t.NumField()
	out := map[string]reflect.StructField{}
	for i := 0; i < total; i++ {
		field := t.Field(i)
		tag := field.Tag.Get("yaml")
		index := strings.Index(tag, ",")
		if index == -1 {
			index = len(tag)
		}
		out[tag[:index]] = field
	}
	return out
}

func getString(in interface{}) string {
	switch s := in.(type) {
	case string:
		return s
	case *string:
		return *s
	}
	return ""
}

func getBool(in interface{}) bool {
	switch b := in.(type) {
	case bool:
		return b
	case *bool:
		return *b
	}
	return false
}

func setElemField(field reflect.Value, value interface{}) {
	switch v := value.(type) {
	case string:
		s := getString(v)
		field.Set(reflect.ValueOf(&s))
	case *string:
		s := getString(v)
		field.Set(reflect.ValueOf(&s))
	case bool:
		b := getBool(v)
		field.Set(reflect.ValueOf(&b))
	case *bool:
		b := getBool(v)
		field.Set(reflect.ValueOf(&b))
	case []string:
		// TODO: Implement []string type
	case map[string]interface{}:
		// TODO: Implement Meta value
	}
}

func mapInterfaceToOptionsModel(in interface{}) *models.EnvironmentItemOptionsModel {
	opt := models.EnvironmentItemOptionsModel{}
	switch m := in.(type) {
	case map[interface{}]interface{}:
		rtype := reflect.TypeOf(opt)
		relem := reflect.ValueOf(&opt).Elem()
		fieldMap := createYamlTagToFieldMap(rtype)

		for k, v := range m {
			field, ok := fieldMap[getString(k)]
			if !ok {
				// TODO: Error handling
				continue
			}
			setElemField(relem.FieldByIndex(field.Index), v)
		}
	}

	return &opt
}

func safeDeref[T any](ptr *T) T {
	if ptr == nil {
		var zero T
		return zero
	}
	return *ptr
}

func gatherOptions(in interface{}) XPutOptions {
	model := *mapInterfaceToOptionsModel(in)
	return XPutOptions{
		Description: safeDeref(model.Description),
		Summary:     safeDeref(model.Summary),
		Title:       safeDeref(model.Title),
		IsRequired:  safeDeref(model.IsRequired),
	}
}

func CollectXPuts(models []models.EnvironmentItemModel) []XPut {
	output := make([]XPut, len(models))
	options := XPutOptions{}

	for index, model := range models {
		var identifier string
		var value any

		for key, value := range model {
			if key == "opts" {
				options = gatherOptions(value)
			} else {
				identifier = key
			}
		}

		output[index] = XPut{
			Identifier: identifier,
			Options:    options,
			Value:      value,
		}
	}

	return output
}
