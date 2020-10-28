package filter

import (
	"reflect"
	"github.com/flosch/pongo2"
	"strings"
)

func init() {
	pongo2.RegisterFilter("ValueWithMap", valueWithMap)
	pongo2.RegisterFilter("HasPrefix", hasPrefix)
	pongo2.RegisterFilter("HasSuffix", hasSuffix)
	pongo2.RegisterFilter("CompareString", compareString)
}

////////////////////////////////////////////////////////////////////////////////
func valueWithMap(in, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	var source = in.Interface()
	var key = param.Interface()

	if source == nil {
		return nil, nil
	}

	if key == nil {
		return nil, nil
	}

	var sourceValue = reflect.ValueOf(source)
	if sourceValue.IsNil() {
		return nil, nil
	}

	switch sourceValue.Kind() {
	case reflect.Map:
		var targetValue = reflect.ValueOf(key)
		if targetValue.IsValid() {
			return pongo2.AsValue(sourceValue.MapIndex(targetValue).Interface()), nil
		}
	}
	return nil, nil
}

////////////////////////////////////////////////////////////////////////////////
func hasPrefix(in, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	return pongo2.AsValue(strings.HasPrefix(in.String(), param.String())), nil
}

func hasSuffix(in, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	return pongo2.AsValue(strings.HasSuffix(in.String(), param.String())), nil
}

func compareString(in, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	return pongo2.AsValue(strings.Compare(in.String(), param.String())), nil
}