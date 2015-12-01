package config

import (
	"os"
	"testing"

	"github.com/zhevron/match"
)

func TestLoadYAML(t *testing.T) {
	yml := "test: value"
	err := LoadYAML(yml)
	_, found := vars["test"]
	match.IsNil(t, err)
	match.Equals(t, found, true)
	yml = `
illegalyaml
test: value
`
	err = LoadYAML(yml)
	match.IsNotNil(t, err)
}

func TestSaveYAML(t *testing.T) {
	yml, err := SaveYAML()
	match.IsNil(t, err)
	match.GreaterThan(t, len(yml), 0)
}

func TestLoadFileYAML(t *testing.T) {
	yml := `
test: value
other:
  nested: 0
`
	f, _ := os.Create(file.Name())
	f.WriteString(yml)
	f.Close()
	err := LoadFileYAML(file.Name())
	_, found := vars["test"]
	match.IsNil(t, err)
	match.Equals(t, found, true)
	err = LoadFileYAML("nonexistant.file")
	match.IsNotNil(t, err)
}

func TestSaveFileYAML(t *testing.T) {
	err := SaveFileYAML(file.Name())
	fi, _ := os.Stat(file.Name())
	match.IsNil(t, err)
	match.NotEquals(t, fi.Size(), 0)
}

func BenchmarkLoadYAML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoadYAML(`
test: value
other:
	nested: 0
		`)
	}
}

func BenchmarkSaveYAML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SaveYAML()
	}
}
