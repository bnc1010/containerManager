package utils

import (
	"fmt"
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

func MapII2MapMap(_map map[string]interface{}) *map[string]map[string]string {
	resMap := map[string]map[string]string {}
	for k, v := range _map {
		tmpMap := map[string]string {}
		for kk, vv := range v.(map[string] interface {}) {
			tmpMap[fmt.Sprintf("%v", kk)] = fmt.Sprintf("%v", vv)
		}
		resMap[k] = tmpMap
	}
	return &resMap
}