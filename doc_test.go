package config

import "fmt"

func ExampleGetString() {
	str := GetString("myVar", "defaultValue")
	fmt.Println(str)
}

func ExampleGetString_nested() {
	str := GetString("myVar.nestedValue", "defaultValue")
	fmt.Println(str)
}

func ExampleLoadFileJSON() {
	if err := LoadFileJSON("/path/to/file.json"); err != nil {
		// Failed to load the configuration file
		panic(err)
	}
}

func ExampleSaveFileJSON() {
	if err := SaveFileJSON("/path/to/file.json"); err != nil {
		// Failed to save the configuration file
		panic(err)
	}
}

func ExampleLoadFileYAML() {
	if err := LoadFileYAML("/path/to/file.yml"); err != nil {
		// Failed to load the configuration file
		panic(err)
	}
}

func ExampleSaveFileYAML() {
	if err := SaveFileYAML("/path/to/file.yml"); err != nil {
		// Failed to save the configuration file
		panic(err)
	}
}
