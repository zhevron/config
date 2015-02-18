package yaml

import "fmt"

func ExampleLoadFile() {
	if err := LoadFile("/path/to/file.yml"); err != nil {
		// Failed to load the configuration file
		panic(err)
	}
}

func ExampleSaveFile() {
	if err := SaveFile("/path/to/file.yml"); err != nil {
		// Failed to save the configuration file
		panic(err)
	}
}

func ExampleGetString() {
	/*
		MyVar: SomeValue
	*/

	str := GetString("MyVar", "DefaultValue")
	fmt.Println(str)
}

func ExampleGetString_nested() {
	/*
		My:
			Nested:
				Var: SomeValue
	*/

	str := GetString("My.Nested.Var", "DefaultValue")
	fmt.Println(str)
}
