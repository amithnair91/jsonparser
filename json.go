package json

import (
	j "encoding/json"
	"fmt"
	"errors"
)

type Type int

const (
	Object  Type = 1
	Array   Type = 2
	String  Type = 3
	Integer Type = 4
)

type JSON struct {
	content interface{}
}

type JSONObject struct {
	content map[string]interface{}
}

type JSONArray struct {
	content []interface{}
}

func NewJSON(jsonBytes []byte) (json JSON, err error) {
	jsonMap, unmarshalError := unmarshallAsJsonObject(jsonBytes)

	if unmarshalError == nil {
		json = JSON{content: jsonMap,}
		return
	}

	jsonArray, unmarshalError := unmarshallAsJsonArray(jsonBytes)

	if unmarshalError == nil {
		json = JSON{content: jsonArray,}
		return
	}

	err = fmt.Errorf("failed to unmarshal json :%v", err)

	return
}

func unmarshallAsJsonObject(jsonBytes []byte) (json interface{}, err error) {
	err = j.Unmarshal(jsonBytes, &json)
	if err != nil {
		err = fmt.Errorf(fmt.Sprintf("failed to unmarshal json :%v", err))
	}
	return
}

func unmarshallAsJsonArray(jsonBytes []byte) (json interface{}, err error) {
	err = j.Unmarshal(jsonBytes, &json)
	if err != nil {
		err = fmt.Errorf(fmt.Sprintf("failed to unmarshal json :%v", err))
	}
	return
}

func (js *JSON) Key(key string) (result interface{}, jsontype Type, err error) {
	fmt.Println(fmt.Sprintf("%#v", js.content))
	jsMap := js.content

	if _, ok := js.content.(map[string]interface{}); ok {
		result = jsMap.(map[string]interface{})[key]
		jsontype = findType(result)
	} else {
		err = errors.New("json node not of type object")
	}

	return result, jsontype, err
}

func (js *JSON) Next(result interface{}) (jsonType Type) {

	return
}

func findType(result interface{}) (jsonType Type) {

	switch result.(type) {
	case string:
		jsonType = String
	case int:
		jsonType = Integer
	case float64:
		jsonType = Integer
	case float32:
		jsonType = Integer
	case map[string]interface{}:
		jsonType = Object
	case []interface{}:
		jsonType = Array
	default:
		jsonType = 0
	}
	return jsonType
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
