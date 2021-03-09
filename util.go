package ujson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ToJsonString(any interface{}) (result string) {
	if any == nil {
		result = ""
	} else {
		bytes, err := json.Marshal(any)
		if err == nil {
			result = string(bytes)
		} else {
			result = ""
		}
	}
	return
}

func ToJsonRawString(any interface{}) (result string) {
	if any == nil {
		result = ""
	} else {
		bytes, err := json.Marshal(any)
		if err == nil {
			result = string(json.RawMessage(bytes))
		} else {
			result = ""
		}
	}
	return
}

func ToJsonByte(any interface{}) (result []byte) {
	if any == nil {
	} else {
		var err error
		result, err = json.Marshal(any)
		if err != nil {
			result = make([]byte, 0)
		}
	}
	return
}

func FromFile(path string, config interface{}) (err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Errorf("Read file error %v.\n", err)
	}
	err = json.Unmarshal(data, config)
	if err != nil {
		fmt.Errorf("Unmarshal error %v.\n", err)
	}
	return
}
