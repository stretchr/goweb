package goweb

import (
	"os"
	"reflect"
	"strconv"
    "url"
    "fmt"
)

// Fill a struct `v` from the values in `form`
func UnmarshalForm(form url.Values, v interface{}) os.Error {
    // check v is valid
    rv := reflect.ValueOf(v).Elem()
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return fmt.Errorf("v must be a valid pointer") 
	}
    // get value
    rv = rv.Elem()
    rt := rv.Type()

    if rv.Kind() == reflect.Struct {
        // for each struct field on v
        for i := 0; i < rt.NumField(); i++ {
            err := unmarshalField(form, rt.Field(i).Name, rv.Field(i))
            if err != nil {
                return err
            }
        }
    } else if rv.Kind() == reflect.Map {
        // for each form value add it to the map
        for k,v := range form {
            if len(v) > 0 {
                rv.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v[0]))
            }
        }
    } else {
        return fmt.Errorf("v must point to a struct or map type")
    }
    return nil
}

func unmarshalField(form url.Values, name string, v reflect.Value) os.Error {
    // form field value
    fvs := form[name]
    if len(fvs)== 0 {
        return nil
    }
    fv := fvs[0]
    // string -> type conversion
    switch v.Kind() {
        case reflect.Int64:
            // convert to Int64
            if i,err := strconv.Atoi64(fv); err == nil {
                v.SetInt(i)
            } 
        case reflect.Int:
            // convert to Int
            // convert to Int64
            if i,err := strconv.Atoi64(fv); err == nil {
                v.SetInt(i)
            } 
        case reflect.String:
            // copy string
            v.SetString(fv)
        case reflect.Bool:
            // the following strings convert to true
            // 1,true,on,yes
            if fv == "1" || fv == "true" || fv == "on" || fv == "yes" {
                v.SetBool(true)
            }
    }
    return nil
} 
