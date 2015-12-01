package config

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
)

// LoadJSON parses the JSON in data to it's internal map.
func LoadJSON(data string) error {
	m := make(stringMap)
	if err := json.Unmarshal([]byte(data), &m); err != nil {
		return err
	}
	mutex.Lock()
	vars = loadMapJSON(m)
	mutex.Unlock()
	return nil
}

// loadMapJSON parses a loaded map structure and adds it to the current
// configuration.
func loadMapJSON(m stringMap) stringMap {
	res := make(stringMap)
	for k, v := range m {
		if reflect.TypeOf(v).Kind() == reflect.Map {
			res[k] = loadMapJSON(v.(map[string]interface{}))
		} else {
			res[k] = v
		}
	}
	return res
}

// SaveJSON returns the current configuration in JSON format.
func SaveJSON() (string, error) {
	mutex.Lock()
	b, err := json.Marshal(&vars)
	mutex.Unlock()
	return string(b), err
}

// LoadFileJSON reads a specified file into memory and parses it using Load().
func LoadFileJSON(path string) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	jsonData := string(b[:])
	return LoadJSON(jsonData)
}

// SaveFileJSON writes the output from Save() to the specified file.
func SaveFileJSON(path string) error {
	data, err := SaveJSON()
	if err == nil {
		err = ioutil.WriteFile(path, []byte(data), 0644)
	}
	return err
}
