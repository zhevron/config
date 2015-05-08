// Copyright (C) 2014-2015 Thomas Lokshall
// Use of this source code is governed by the MIT license.
// See LICENSE.md for details.

// Package json implements functions to read and write JSON configuration files.
package json

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"strings"
	"sync"
)

// stringMap is a shorthand type for the map that holds the variables.
type stringMap map[string]interface{}

// vars holds the variables currently loaded into the package.
var vars stringMap

// mutex holds the sync.Mutex used to control concurrent reads/writes.
var mutex *sync.Mutex

// Initialize the vars map on init.
func init() {
	vars = make(stringMap)
	mutex = &sync.Mutex{}
}

// Load parses the JSON in data to it's internal map.
func Load(data string) error {
	m := make(stringMap)
	if err := json.Unmarshal([]byte(data), &m); err != nil {
		return err
	}
	mutex.Lock()
	vars = loadMap(m)
	mutex.Unlock()
	return nil
}

// loadMap parses a loaded map structure and adds it to the current
// configuration.
func loadMap(m map[string]interface{}) stringMap {
	res := make(stringMap)
	for k, v := range m {
		if reflect.TypeOf(v).Kind() == reflect.Map {
			res[k] = loadMap(v.(map[string]interface{}))
		} else {
			res[k] = v
		}
	}
	return res
}

// Save returns the current configuration in JSON format.
func Save() (string, error) {
	mutex.Lock()
	b, err := json.Marshal(&vars)
	mutex.Unlock()
	return string(b), err
}

// LoadFile reads a specified file into memory and parses it using Load().
func LoadFile(path string) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	jsonData := string(b[:])
	return Load(jsonData)
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
	mutex.Lock()
	vars[key] = val
	mutex.Unlock()
}

// Remove deletes the value with the given key from the configuration.
func Remove(key string) {
	mutex.Lock()
	delete(vars, key)
	mutex.Unlock()
}

// Get returns the value of the given key from the current configuration or the
// value of def if not found.
// The value is returned as an interface{}.
func Get(key string, def interface{}) interface{} {
	mutex.Lock()
	v := getFromMap(vars, key, def)
	mutex.Unlock()
	return v
}

// getFromMap returns the value of the given key from the specified map or the
// value of def if not found.
// The value is returned as an interface{}.
// If the key contains a dot (.), the function will be called recursively.
func getFromMap(m stringMap, key string, def interface{}) interface{} {
	if n := strings.Index(key, "."); n != -1 {
		if _, found := m[key[:n]]; found {
			if _, ok := m[key[:n]].(stringMap); ok {
				return getFromMap(m[key[:n]].(stringMap), key[n+1:], def)
			}
		}
		return def
	}
	if v, found := m[key]; found {
		return v
	}
	return def
}

// GetInt returns the value of the given key from the current configuration
// or the value of def if not found.
// The value is returns as an int.
// For example usage, see the GetString examples.
func GetInt(key string, def int) int {
	return Get(key, def).(int)
}

// GetInt8 returns the value of the given key from the current configuration
// or the value of def if not found.
// The value is returns as an int8.
// For example usage, see the GetString examples.
func GetInt8(key string, def int8) int8 {
	return int8(Get(key, def).(int))
}

// GetInt16 returns the value of the given key from the current configuration
// or the value of def if not found.
// The value is returns as an int16.
// For example usage, see the GetString examples.
func GetInt16(key string, def int16) int16 {
	return int16(Get(key, def).(int))
}

// GetInt32 returns the value of the given key from the current configuration
// or the value of def if not found.
// The value is returns as an int32.
// For example usage, see the GetString examples.
func GetInt32(key string, def int32) int32 {
	return int32(Get(key, def).(int))
}

// GetInt64 returns the value of the given key from the current configuration
// or the value of def if not found.
// The value is returns as an int64.
// For example usage, see the GetString examples.
func GetInt64(key string, def int64) int64 {
	return int64(Get(key, def).(int))
}

// GetUint returns the value of the given key from the current configuration
// or the value of def if not found.
// The value is returns as an uint.
// For example usage, see the GetString examples.
func GetUint(key string, def uint) uint {
	return uint(Get(key, def).(int))
}

// GetUint8 returns the value of the given key from the current configuration
// or the value of def if not found.
// The value is returns as an uint8.
// For example usage, see the GetString examples.
func GetUint8(key string, def uint8) uint8 {
	return uint8(Get(key, def).(int))
}

// GetUint16 returns the value of the given key from the current configuration
// or the value of def if not found.
// The value is returns as an uint16.
// For example usage, see the GetString examples.
func GetUint16(key string, def uint16) uint16 {
	return uint16(Get(key, def).(int))
}

// GetUint32 returns the value of the given key from the current configuration
// or the value of def if not found.
// The value is returns as an uint32.
// For example usage, see the GetString examples.
func GetUint32(key string, def uint32) uint32 {
	return uint32(Get(key, def).(int))
}

// GetUint64 returns the value of the given key from the current configuration
// or the value of def if not found.
// The value is returns as an uint64.
// For example usage, see the GetString examples.
func GetUint64(key string, def uint64) uint64 {
	return uint64(Get(key, def).(int))
}

// GetFloat32 returns the value of the given key from the current configuration
// or the value of def if not found.
// The value is returns as a float32.
// For example usage, see the GetString examples.
func GetFloat32(key string, def float32) float32 {
	return float32(Get(key, def).(float64))
}

// GetFloat64 returns the value of the given key from the current configuration
// or the value of def if not found.
// The value is returns as a float64.
// For example usage, see the GetString examples.
func GetFloat64(key string, def float64) float64 {
	return Get(key, def).(float64)
}

// GetBool returns the value of the given key from the current configuration
// or the value of def if not found.
// The value is returns as a bool.
// For example usage, see the GetString examples.
func GetBool(key string, def bool) bool {
	return Get(key, def).(bool)
}

// GetString returns the value of the given key from the current configuration
// or the value of def if not found.
// The value is returns as a string.
func GetString(key string, def string) string {
	return Get(key, def).(string)
}
