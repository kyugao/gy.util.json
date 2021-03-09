package ujson

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	data := `{"Name":"Happy"}`
	t.Logf("Config = %s.", ToJsonRawString(data))
}
