package ujson

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	data := `{"Name":"Happy"}`
	t.Logf("Config = %s.", ToJsonRawString(data))
}

var sample = `{"newPassword": "MTIzNDU2", "oldPassword": "MTIzNDU2"}`

func TestToJsonString(t *testing.T) {
	t.Log(ToJsonRawString(sample))
}