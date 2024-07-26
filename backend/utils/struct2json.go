package tools

import "encoding/json"

// Tojsonify make Struct to jsonify, can be inserted into LocalDatabase.
func Tojsonify(anyInterface interface{}) string {
	getMapper := anyInterface
	getMasher, err := json.Marshal(getMapper)
	if err != nil {
		return ""
	}
	return string(getMasher)
}

// JsonToInterface Make Json to map[] == >
func JsonToInterface(anyJson string) map[string]interface{} {
	var JsonUnmashell map[string]interface{}
	json.Unmarshal([]byte(anyJson), &JsonUnmashell)
	return JsonUnmashell
}

func Listize(anyJson string) []interface{} {
	var JsonUnmashell []interface{}
	json.Unmarshal([]byte(anyJson), &JsonUnmashell)
	return JsonUnmashell
}

func StructPrinter(anyStruct interface{}) {
	getMapper := anyStruct
	getMasher, _ := json.Marshal(getMapper)
	print(string(getMasher))
}
