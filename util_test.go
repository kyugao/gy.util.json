package ujson

import (
	"testing"
)

var Config struct {
	Item1 string
	Item2 int
	Item3 bool
}

func TestLoadConfig(t *testing.T) {
	FromFile("./config.json", &Config)
	t.Logf("Config = %s.\n", ToJsonRawString(&Config))
}
