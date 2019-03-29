package json

import (
	json "github.com/json-iterator/go"
)

var Json = json.ConfigCompatibleWithStandardLibrary

func Marshal(v interface{}) ([]byte, error) {
	return Json.Marshal(v)
}

func MarshalToString(v interface{}) (string, error) {
	return Json.MarshalToString(v)
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return Json.MarshalIndent(v, prefix, indent)
}

func UnmarshalFromString(str string, v interface{}) error {
	return Json.UnmarshalFromString(str, v)
}

func Unmarshal(data []byte, v interface{}) error {
	return Json.Unmarshal(data, v)
}

func Valid(data []byte) bool {
	return Json.Valid(data)
}
