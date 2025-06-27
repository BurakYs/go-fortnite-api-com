package fortniteapi

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func structToQuery(params any, defaultLanguage Language) (url.Values, error) {
	query := url.Values{}
	if params == nil {
		return query, nil
	}

	value := reflect.ValueOf(params)
	typ := reflect.TypeOf(params)

	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("params must be a struct, got %s", typ.Kind())
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		key := field.Tag.Get("url")
		if key == "" {
			key = field.Name
		} else {
			key = strings.Split(key, ",")[0]
		}

		fieldValue := value.Field(i)
		if fieldValue.IsZero() {
			if key == "language" {
				query.Set(key, string(defaultLanguage))
			}

			continue
		}

		query.Set(key, fmt.Sprint(fieldValue.Interface()))
	}

	return query, nil
}
