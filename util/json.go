package util

import (
	"bytes"
	"encoding/json"
	"strings"
)

type SafeToJsonOption struct {
	DoEscapeHTML bool
}

type SafeToJsonOptionFn func(option *SafeToJsonOption)

func SafeToJson(o interface{}, optionFns ...SafeToJsonOptionFn) string {
	option := new(SafeToJsonOption)
	for _, optionFn := range optionFns {
		optionFn(option)
	}

	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(option.DoEscapeHTML)
	err := jsonEncoder.Encode(o)
	if err != nil {
		return ""
	}
	return strings.TrimRight(bf.String(), "\n")
}

// UnmarshalToMap Note: can not use json.Unmarshal, uint64 will lose accuracy
func UnmarshalToMap(jsonBytes []byte) (map[string]interface{}, error) {
	var objForUpdateMap map[string]interface{}
	d := json.NewDecoder(bytes.NewReader(jsonBytes))
	d.UseNumber()
	err := d.Decode(&objForUpdateMap)
	if err != nil {
		return nil, err
	}
	return objForUpdateMap, nil
}

func SafeUnmarshalToMap(jsonBytes []byte) map[string]interface{} {
	resMap, _ := UnmarshalToMap(jsonBytes)
	return resMap
}

func SafeUnmarshalStructToMap(obj interface{}) map[string]interface{} {
	jsonStr := SafeToJson(obj)
	return SafeUnmarshalToMap(String2Bytes(&jsonStr))
}
