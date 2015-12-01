package config

import (
	"io/ioutil"
	"reflect"

	"gopkg.in/yaml.v2"
)

// LoadYAML parses the YAML in data to it's internal map.
func LoadYAML(data string) error {
	m := make(interfaceMap)
	if err := yaml.Unmarshal([]byte(data), &m); err != nil {
		return err
	}
	mutex.Lock()
	vars = loadMapYAML(m)
	mutex.Unlock()
	return nil
}

// loadMapYAML parses a loaded map structure and adds it to the current
// configuration.
// The prefix will be added before all values.
func loadMapYAML(m interfaceMap) stringMap {
	res := make(stringMap)
	for k, v := range m {
		if reflect.TypeOf(v).Kind() == reflect.Map {
			res[k.(string)] = loadMapYAML(v.(interfaceMap))
		} else {
			res[k.(string)] = v
		}
	}
	return res
}

// SaveYAML returns the current configuration in YAML format.
func SaveYAML() (string, error) {
	mutex.Lock()
	b, err := yaml.Marshal(&vars)
	mutex.Unlock()
	return string(b), err
}

// LoadFileYAML reads a specified file into memory and parses it using Load().
func LoadFileYAML(path string) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	yml := string(b[:])
	return LoadYAML(yml)
}

// SaveFileYAML writes the output from Save() to the specified file.
func SaveFileYAML(path string) error {
	data, err := SaveYAML()
	if err == nil {
		err = ioutil.WriteFile(path, []byte(data), 0644)
	}
	return err
}
