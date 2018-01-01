package json_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	j "jsonparser"
	"fmt"
)

const InvalidJson = `invalidJSon`
const SimpleJson = `{"name":"json"}`

func TestBuildJsonFailsForInvalidJson(t *testing.T) {
	jsonStr := []byte(InvalidJson)

	_, err := j.NewJSON(jsonStr)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to unmarshal json :")
}

func TestBuildJsonReturnsJSONForValidJson(t *testing.T) {
	jsonBytes := []byte(SimpleJson)
	expectedType := j.JSON{}

	json, err := j.NewJSON(jsonBytes)

	assert.NoError(t, err)
	assert.NotNil(t, json)
	assert.IsType(t, expectedType, json)
}

func TestKeyReturnsNilIfKeyDoesNotExist(t *testing.T) {
	jsonBytes := []byte(SimpleJson)
	json, _ := j.NewJSON(jsonBytes)

	result, jsonType := json.Key("non_existent_key")

	assert.Nil(t, result)
	assert.Equal(t, j.JsonType(0), jsonType)
}

func TestKeyReturnsValueWithTypeJSONObject(t *testing.T) {
	jsonObject := `{"firstlevel":"jsonObjectValue"}`
	jsonWithJsonObject := fmt.Sprintf(`{"key":%s}`, jsonObject)
	expectedResult := map[string]interface{}{
		"firstlevel": "jsonObjectValue",
	}
	jsonBytes := []byte(jsonWithJsonObject)

	json, _ := j.NewJSON(jsonBytes)
	result, jsonType := json.Key("key")

	assert.NotNil(t, result)
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, j.JsonType(1), jsonType)
}

func TestKeyReturnsValueWithTypeJSONArray(t *testing.T) {
	jsonArray := `["1","2","3","4"]`
	jsonWithJsonArray := fmt.Sprintf(`{"key":%s}`, jsonArray)
	expectedResult := []interface{}{
		"1", "2", "3", "4",
	}
	jsonBytes := []byte(jsonWithJsonArray)

	json, _ := j.NewJSON(jsonBytes)
	result, jsonType := json.Key("key")

	assert.NotNil(t, result)
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, j.JsonType(2), jsonType)
}

func TestKeyReturnsValueWithTypeString(t *testing.T) {
	value := `json value`
	jsonWithStringValue := fmt.Sprintf(`{"key":"%s"}`, value)
	jsonBytes := []byte(jsonWithStringValue)

	json, _ := j.NewJSON(jsonBytes)
	result, jsonType := json.Key("key")

	assert.NotNil(t, result)
	assert.Equal(t, value, result)
	assert.Equal(t, j.JsonType(3), jsonType)
}

func TestKeyReturnsValueWithTypeInteger(t *testing.T) {
	jsonWithIntegerValue := `{"key":123}`
	expectedResult := float64(123)
	jsonBytes := []byte(jsonWithIntegerValue)

	json, _ := j.NewJSON(jsonBytes)
	result, jsonType := json.Key("key")

	assert.NotNil(t, result)
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, j.JsonType(4), jsonType)
}
