package config

import (
	"os"
	"reflect"
	"testing"

	"github.com/Thomasdezeeuw/ini"
)

func TestConfigParseFile_IncorrectFilePath(t *testing.T) {
	err := New().ParseFile("file_that_does_not_exist", ReadFromFile)
	if err == nil {
		t.Errorf(msg, "not found error", nil)
	}
}

func TestConfigParseFile_IncorrectConfigFile(t *testing.T) {
	err := New().ParseFile("./testdata/incorrect.ini", ReadFromFile)
	if err == nil {
		t.Errorf(msg, "incorrect format error", nil)
	}
}

func TestConfigParseFile_ExtendConfig(t *testing.T) {
	exp := map[string]map[string]string{
		ini.Global: {
			"name": "test",
			"test": "yes",
			"base": "yes",
		},
		"test": {
			"a": "b",
			"x": "z",
		},
		"smth": {
			"hello": "world",
		},
		systemSection: {
			keyActiveSection: "test",
			keyExtend:        "base.ini",
		},
	}
	c := New()
	err := c.ParseFile("./testdata/test.ini", ReadFromFile)
	if err != nil {
		t.Errorf(msg, nil, err)
	}
	if !reflect.DeepEqual(c.data, exp) {
		t.Errorf(msg, exp, c.data)
	}
}

func TestConfigParseFile_Cycle(t *testing.T) {
	err := New().ParseFile("./testdata/loop.ini", ReadFromFile)
	if err == nil {
		t.Errorf(msg, "cycle inheritance not allowed", nil)
	}
}

func TestConfigStringDefault(t *testing.T) {
	c := New()
	c.ParseFile("./testdata/test.ini", ReadFromFile)
	if r := c.StringDefault("x", "some default"); r != "z" {
		t.Errorf(msg, "z", r)
	}
	exp := "some default"
	if r := c.StringDefault("key_that_does_not_exist", exp); r != exp {
		t.Errorf(msg, exp, r)
	}
	exp = "test"
	if r := c.StringDefault("name", "xxx"); r != exp {
		t.Errorf(msg, exp, r)
	}

	c = New()
	c.ParseFile("./testdata/test.ini", "test")
	if r := c.StringDefault("x", "some default"); r != "z" {
		t.Errorf(msg, "z", r)
	}
}

func TestConfigGetString_EnvVar(t *testing.T) {
	c := New()
	c.data = map[string]map[string]string{
		"": {
			"test": "${GOPATH}",
		},
	}
	exp := os.Getenv("GOPATH")
	if r, _ := c.getString("", "test"); r != exp || r == "" {
		t.Errorf(msg, exp, r)
	}
}

func TestConfigSection(t *testing.T) {
	c := New()
	c.ParseFile("./testdata/test.ini", ReadFromFile)
	exp := map[string]string{
		"a": "b",
		"x": "z",
	}
	if r := c.Section("test"); !reflect.DeepEqual(r, exp) {
		t.Errorf(msg, exp, r)
	}
}

func TestRelFilepath(t *testing.T) {
	exp := "/home/user/path/to/dir/file.ini"
	if r := relFilepath("./config/app.ini", exp); r != exp {
		t.Errorf(msg, exp, r)
	}
}

var msg = `Incorrect result. Expected "%v", got "%v".`
