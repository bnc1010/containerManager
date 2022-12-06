package utils


import (
	"fmt"
)


func MapInterface2String(inputData map[string]interface{}) map[string]string {
    outputData := map[string]string{}
    for key, value := range inputData {
        switch value.(type) {
        case string:
            outputData[key] = value.(string)
        case int:
            tmp := value.(int)
            outputData[key] = fmt.Sprintf("%d", tmp)
        case int64:
            tmp := value.(int64)
            outputData[key] = fmt.Sprintf("%d", tmp)
        }
    }
    return outputData
}

func GetInterfaceToInt32(t1 interface{}) int32 {
	var t2 int32
	switch t1.(type) {
	case uint:
		t2 = int32(t1.(uint))
		break
	case int8:
		t2 = int32(t1.(int8))
		break
	case uint8:
		t2 = int32(t1.(uint8))
		break
	case int16:
		t2 = int32(t1.(int16))
		break
	case uint16:
		t2 = int32(t1.(uint16))
		break
	case int32:
		t2 = t1.(int32)
		break
	case uint32:
		t2 = int32(t1.(uint32))
		break
	case float32:
		t2 = int32(t1.(float32))
		break
	case float64:
		t2 = int32(t1.(float64))
		break
	default:
		t2 = t1.(int32)
		break
	}
	return t2
}


func GetInterfaceToInt(t1 interface{}) int {
	return int(GetInterfaceToInt32(t1))
}
