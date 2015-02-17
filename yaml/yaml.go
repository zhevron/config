// Copyright (C) 2014-2015 Thomas Lokshall
// Use of this source code is governed by the MIT license.
// See LICENSE.md for details.

// Package yaml implements a way to easily read and write YAML configuration files.
//
// To read a configuration file, simply call the LoadFile method on a new Configuration object:
//		import cfg "gopkg.in/zhevron/gov1/yaml"
//		if err := cfg.LoadFile("/path/to/file.yml"); err != nil {
//			// Failed to load the configuration file
//		}
//
// To save the current configuration to a file, call the SaveFile method:
//		if err := cfg.SaveFile("/path/to/file.yml"); err != nil {
//			// Failed to save the configuration file
//		}
//
// You can easily access vars with the Get method or any of the Get_Type_ methods:
//		host := cfg.GetString("host", "localhost")
//		port := cfg.GetInt("port", 10000)
//
// You can just as easily set a value using the Set method:
//		cfg.Set("host", host)
//		cfg.Set("port", port)
//
// Nested objects are accessed by adding a dot (.) between the value names.
// If you have a YAML file that looks like:
//		some:
//			value: 1
// The value can be accessed by calling the Get method:
//		val := cfg.GetInt("some.value", 1)
package yaml

import (
	"io/ioutil"
	"reflect"
	"strings"

	"gopkg.in/yaml.v1"
)

// interfaceMap is a shorthand type required by the YAML library.
type interfaceMap map[interface{}]interface{}

// stringMap is a shorthand type for the map that holds the variables.
type stringMap map[string]interface{}

// vars holds the variables currently loaded into the package.
var vars stringMap

// Initialize the vars map on init.
func init() {
	vars = make(stringMap)
}

// Load parses the YAML in data to it's internal map.
func Load(data string) error {
	m := make(interfaceMap)
	if err := yaml.Unmarshal([]byte(data), &m); err != nil {
		return err
	}
	vars = loadMap(m)
	return nil
}

// loadMap parses a loaded map structure and adds it to the current configuration.
// The prefix will be added before all values.
func loadMap(m map[interface{}]interface{}) stringMap {
	res := make(stringMap)
	for k, v := range m {
		if reflect.TypeOf(v).Kind() == reflect.Map {
			res[k.(string)] = loadMap(v.(map[interface{}]interface{}))
		} else {
			res[k.(string)] = v
		}
	}
	return res
}

// Save returns the current configuration in YAML format.
func Save() (string, error) {
	b, err := yaml.Marshal(&vars)
	return string(b), err
}

// LoadFile reads a specified file into memory and parses it using Load().
func LoadFile(path string) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	yml := string(b[:])
	return Load(yml)
}

// SaveFile writes the output from Save() to the specified file.
func SaveFile(path string) error {
	data, err := Save()
	if err == nil {
		err = ioutil.WriteFile(path, []byte(data), 0644)
	}
	return err
}

// Set sets a given value in the configuration to an arbitrary type.
func Set(key string, val interface{}) {
	vars[key] = val
}

// Remove deletes the value with the given key from the configuration.
func Remove(key string) {
	delete(vars, key)
}

// Get returns the value of the given key from the current configuration or the value of def if not found.
// The value is returned as an interface{}.
func Get(key string, def interface{}) interface{} {
	return getFromMap(vars, key, def)
}

// getFromMap returns the value of the given key from the specified map or the value of def if not found.
// The value is returned as an interface{}.
// If the key contains a dot (.), the function will be called recursively.
func getFromMap(m stringMap, key string, def interface{}) interface{} {
	if n := strings.Index(key, "."); n != -1 {
		return getFromMap(m[key[:n]].(stringMap), key[n+1:], def)
	}
	if v, found := m[key]; found {
		return v
	}
	return def
}

// GetInt returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as an int.
func GetInt(key string, def int) int {
	return Get(key, def).(int)
}

// GetInt8 returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as an int8.
func GetInt8(key string, def int8) int8 {
	return int8(Get(key, def).(int))
}

// GetInt16 returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as an int16.
func GetInt16(key string, def int16) int16 {
	return int16(Get(key, def).(int))
}

// GetInt32 returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as an int32.
func GetInt32(key string, def int32) int32 {
	return int32(Get(key, def).(int))
}

// GetInt64 returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as an int64.
func GetInt64(key string, def int64) int64 {
	return int64(Get(key, def).(int))
}

// GetUint returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as an uint.
func GetUint(key string, def uint) uint {
	return uint(Get(key, def).(int))
}

// GetUint8 returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as an uint8.
func GetUint8(key string, def uint8) uint8 {
	return uint8(Get(key, def).(int))
}

// GetUint16 returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as an uint16.
func GetUint16(key string, def uint16) uint16 {
	return uint16(Get(key, def).(int))
}

// GetUint32 returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as an uint32.
func GetUint32(key string, def uint32) uint32 {
	return uint32(Get(key, def).(int))
}

// GetUint64 returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as an uint64.
func GetUint64(key string, def uint64) uint64 {
	return uint64(Get(key, def).(int))
}

// GetFloat32 returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as a float32.
func GetFloat32(key string, def float32) float32 {
	return float32(Get(key, def).(float64))
}

// GetFloat64 returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as a float64.
func GetFloat64(key string, def float64) float64 {
	return Get(key, def).(float64)
}

// GetBool returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as a bool.
func GetBool(key string, def bool) bool {
	return Get(key, def).(bool)
}

// GetString returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as a string.
func GetString(key string, def string) string {
	return Get(key, def).(string)
}
