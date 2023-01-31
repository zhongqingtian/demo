package instruct

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"net/url"
	"reflect"
)

type MapStruct struct {
	Name string  `json:"name"`
	Age  int64   `json:"age"`
	F64  float64 `json:"f64"`
	//Ll    []int64   `url:"ll"`
	Times int  `json:"times"`
	User  User `json:"user"`
}

type User struct {
	Name2 string `json:"name_2"`
	Age2  int64  `json:"age_2"`
}

func TurnUrls(a interface{}) url.Values {
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	//  m := make(map[string]string)
	urls := url.Values{}
	for i := 0; i < t.NumField(); i++ {
		var key, val string
		f := t.Field(i)
		name := f.Name
		key = f.Tag.Get("json")
		// fmt.Println("tag:", key)
		// fmt.Println(v.FieldByName(f.Name).Elem().)
		switch f.Type.Kind() {
		case reflect.String:
			val = v.FieldByName(name).String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint,
			reflect.Uint8,
			reflect.Uint16,
			reflect.Uint32,
			reflect.Uint64:
			val = fmt.Sprintf("%d", v.FieldByName(name).Int())
		case reflect.Float32, reflect.Float64:
			val = fmt.Sprintf("%f", v.FieldByName(name).Float())
		default:
			fmt.Println(f.Type.String() + " unsported")
		}
		if val == "" {
			continue
		}
		// fmt.Println(f.Name + "=" + val)
		//  m[key] = val
		urls.Add(key, val)
	}

	return urls
}

func GooleTurnUrls(o interface{}) url.Values {
	// github.com/google/go-querystring/query

	vals, _ := query.Values(o)
	fmt.Println(vals.Encode())
	return vals
}

func ConvertToMap(model interface{}) map[string]string {
	ret := make(map[string]string)

	modelReflect := reflect.ValueOf(model)

	if modelReflect.Kind() == reflect.Ptr {
		modelReflect = modelReflect.Elem()
	}

	modelRefType := modelReflect.Type()
	fieldsCount := modelReflect.NumField()

	var fieldData interface{}
	for i := 0; i < fieldsCount; i++ {
		field := modelReflect.Field(i)

		switch field.Kind() {
		case reflect.Struct:
			fallthrough
		case reflect.Ptr:
			fieldData = ConvertToMap(field.Interface())
		default:
			fieldData = field.Interface()
		}

		ret[modelRefType.Field(i).Tag.Get("json")] = fmt.Sprintf("%v", fieldData)
	}
	return ret
}
