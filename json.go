package json

import (
	j "encoding/json"
	"fmt"
)

type Type int

const (
	Object  Type = 1
	Array   Type = 2
	String  Type = 3
	Integer Type = 4
)

type JSON struct {
	content map[string]interface{}
}

type JSONObject struct {
	content map[string]interface{}
}

type JSONArray struct {
	content []interface{}
}

func NewJSON(jsonBytes []byte) (json JSON, err error) {
	var jsonMap map[string]interface{}
	err = j.Unmarshal(jsonBytes, &jsonMap)
	fmt.Println(fmt.Sprintf("%v", err))
	if err != nil {
		err = fmt.Errorf(fmt.Sprintf("failed to unmarshal json :%v", err))
		return
	}
	json = JSON{content: jsonMap,}
	return
}

func (js *JSON) Key(key string) (result interface{}, jsontype Type) {
	jsMap := js.content
	result = jsMap[key]
	switch result.(type) {
	case string:
		jsontype = String
	case int:
		jsontype = Integer
	case float64:
		jsontype = Integer
	case float32:
		jsontype = Integer
	case map[string]interface{}:
		jsontype = Object
	case []interface{}:
		jsontype = Array
	default:
		jsontype = 0
	}
	return result, jsontype
}

func checkIfObject(str interface{}) (isObject bool) {
	if _, ok := str.(map[string]interface{}); ok {
		isObject = true
	}
	return
}

func checkIfArray(str interface{}) (isArray bool) {
	if _, ok := str.([]interface{}); ok {
		isArray = true
	}
	return
}
