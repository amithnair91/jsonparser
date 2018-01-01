package json

import (
	j "encoding/json"
	"fmt"
)

type JsonType int

const (
	Object  JsonType = 1
	Array   JsonType = 2
	String  JsonType = 3
	Integer JsonType = 4
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

func (js *JSON) Key(key string) (result interface{}, jsontype JsonType) {
	jsMap := js.content
	result = jsMap[key]
	switch result.(type) {
	case string:
		isObject := checkIfObject(result)
		if isObject {
			jsontype = Object
		} else {
			jsontype = String
		}
	case int:
		jsontype = Array
		isArray := checkIfArray(result)
		if isArray {
			jsontype = Array
		} else {
			jsontype = Integer
		}
	default:
		jsontype = 0
	}
	return result,jsontype
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