package run

import (
	"reflect"
	"testing"
)

func TestParseFile_IncorrectFile(t *testing.T) {
	m, err := parseFile("file_that_does_not_exist")
	if m != nil || err == nil {
		t.Errorf("File cannot be open, error expected.")
	}
}

func TestParseFile_NotMap(t *testing.T) {
	m, err := parseFile("testdata/configs/not_map.yml")
	if m != nil || err == nil {
		t.Errorf("YAML file must contain a map, an error expected.")
	}
}

func TestParseFile(t *testing.T) {
	r, err := parseFile("testdata/configs/correct_config.yml")
	assertNil(t, err)
	m, err := parseMap(r, "watch")
	assertNil(t, err)
	exp := map[string][]string{
		"smth":      {"/pass", "/start smth", "goal help test"},
		"smth_else": {"/pass", "/single smth_else"},
	}
	if !reflect.DeepEqual(m, exp) {
		t.Errorf(`Expected %#v, got %#v.`, exp, m)
	}
}

func TestParseMap_IncorrectKey(t *testing.T) {
	r, err := parseFile("testdata/configs/correct_config.yml")
	assertNil(t, err)
	m, err := parseMap(r, "incorrect_key")
	if m != nil || err == nil {
		t.Errorf("Key does not exist, error expected.")
	}
}

func TestParseMap_NotMap(t *testing.T) {
	r, err := parseFile("testdata/configs/correct_config.yml")
	assertNil(t, err)
	m, err := parseMap(r, initSection)
	if m != nil || err == nil {
		t.Errorf(`Section "%s" is not a map, error expected.`, initSection)
	}
}

func TestParseMap_NotSliceMap(t *testing.T) {
	r, err := parseFile("testdata/configs/incorrect_watch.yml")
	assertNil(t, err)
	m, err := parseMap(r, watchSection)
	if m != nil || err == nil {
		t.Errorf(`Section "%s" is not a map[string][]string, error expected.`, watchSection)
	}
}

func TestParseSlice_IncorrectKey(t *testing.T) {
	r, err := parseFile("testdata/configs/correct_config.yml")
	assertNil(t, err)
	m, err := parseSlice(r, "incorrect_key")
	if m != nil || err == nil {
		t.Errorf("Key does not exist, error expected.")
	}
}

func TestParseSlice_NotMap(t *testing.T) {
	r, err := parseFile("testdata/configs/correct_config.yml")
	assertNil(t, err)
	m, err := parseSlice(r, watchSection)
	if m != nil || err == nil {
		t.Errorf(`Section "%s" is not a map, error expected.`, watchSection)
	}
}

func assertNil(t *testing.T, err error) {
	if err != nil {
		t.Errorf("No error expected, got: %v.", err)
	}
}
