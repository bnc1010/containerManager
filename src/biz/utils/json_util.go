package utils

import (
	"encoding/json"
)

func Map2Bytes(_map map[string]interface{}) ([]byte, error) {
	bytes, err := json.Marshal(_map)
    if err != nil {
        return nil, err
    }
	return bytes, nil
}

func Bytes2Map(bytes []byte) (*map[string]interface{}, error) {
	_map := make(map[string]interface{})
	err := json.Unmarshal(bytes, &_map)
    if err != nil {
        return nil, err
    }
	return &_map, nil
}

func Array2Bytes(arr []interface{}) ([]byte, error) {
	bytes, err := json.Marshal(arr)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}


func Bytes2Array(bytes []byte) (*[]interface{}, error) {
	var arr []interface{}
	err := json.Unmarshal(bytes, &arr)
	if err != nil {
		return nil, err
	}
	return &arr, nil
}