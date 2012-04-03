package goweb

import (
	"reflect"
	"strconv"
	"net/url"
	"fmt"
)

// Fill a struct `v` from the values in `form`
func UnmarshalForm(form url.Values, v interface{}) error {
	// check v is valid
	rv := reflect.ValueOf(v).Elem()
	// dereference pointer
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	// get type
	rt := rv.Type()

	if rv.Kind() == reflect.Struct {
		// for each struct field on v
		for i := 0; i < rt.NumField(); i++ {
			err := unmarshalField(form, rt.Field(i), rv.Field(i))
			if err != nil {
				return err
			}
		}
	} else if rv.Kind() == reflect.Map && !rv.IsNil() {
		// for each form value add it to the map
		for k, v := range form {
			if len(v) > 0 {
				rv.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v[0]))
			}
		}
	} else {
		return fmt.Errorf("v must point to a struct or a non-nil map type")
	}
	return nil
}

func unmarshalField(form url.Values, t reflect.StructField, v reflect.Value) error {
	// form field value
	fvs := form[t.Name]
	if len(fvs) == 0 {
		return nil
	}
	fv := fvs[0]
	// string -> type conversion
	switch v.Kind() {
	case reflect.Int64:
		// convert to Int64
		if i, err := strconv.ParseInt(fv, 10, 64); err == nil {
			v.SetInt(i)
		}
	case reflect.Int:
		// convert to Int
		// convert to Int64
		if i, err := strconv.ParseInt(fv, 10, 64); err == nil {
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
	case reflect.Slice:
		// ONLY STRING SLICES SO FAR
		// add all form values to slice
		sv := reflect.MakeSlice(t.Type, len(fvs), len(fvs))
		for i, fv := range fvs {
			svv := sv.Index(i)
			svv.SetString(fv)
		}
		v.Set(sv)
	default:
		fmt.Println("unknown type", v.Kind())
	}
	return nil
}
