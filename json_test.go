package json_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	j "jsonparser"
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


