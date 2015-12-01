package config

import (
	"os"
	"testing"

	"github.com/zhevron/match"
)

func TestLoadJSON(t *testing.T) {
	json := `{"test": "value"}`
	err := LoadJSON(json)
	_, found := vars["test"]
	match.IsNil(t, err)
	match.Equals(t, found, true)
	json = `{illegaljson, "test": "value"}`
	err = LoadJSON(json)
	match.IsNotNil(t, err)
}

func TestSaveJSON(t *testing.T) {
	json, err := SaveJSON()
	match.IsNil(t, err)
	match.NotEquals(t, len(json), 0)
}

func TestLoadFileJSON(t *testing.T) {
	json := `{"test": "value", "other": {"nested": 0}}`
	f, _ := os.Create(file.Name())
	f.WriteString(json)
	f.Close()
	err := LoadFileJSON(file.Name())
	_, found := vars["test"]
	match.IsNil(t, err)
	match.Equals(t, found, true)
	err = LoadFileJSON("nonexistant.file")
	match.IsNotNil(t, err)
}

func TestSaveFileJSON(t *testing.T) {
	err := SaveFileJSON(file.Name())
	fi, _ := os.Stat(file.Name())
	match.IsNil(t, err)
	match.NotEquals(t, fi.Size(), 0)
}

func BenchmarkLoadJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoadJSON(`{"test": "value", "other": {"nested": 0}}`)
	}
}

func BenchmarkSaveJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SaveJSON()
	}
}
