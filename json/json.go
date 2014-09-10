// Copyright (C) 2014 Thomas Lokshall
// Use of this source code is governed by the MIT license.
// See LICENSE.md for details.

// Package json implements a way to easily read and write JSON configuration files.
//
// To read a configuration file, simply call the LoadFile method on a new Configuration object:
//		import "gopkg.in/zhevron/gocfg.v1/json"
//		cfg := json.NewConfiguration()
//		if err := cfg.LoadFile("/path/to/file.json"); err != nil {
//			// Failed to load the configuration file
//		}
//
// To save the current configuration to a file, call the SaveFile method:
//		if err := cfg.SaveFile("/path/to/file.json"); err != nil {
//			// Failed to save the configuration file
//		}
//
// You can easily access values with the Get method or any of the Get_Type_ methods:
//		host := cfg.GetString("host", "localhost")
//		port := cfg.GetInt("port", 10000)
//
// You can just as easily set a value using the Set method:
//		cfg.Set("host", host)
//		cfg.Set("port", port)
//
// Nested objects are accessed by adding a dot (.) between the value names.
// If you have a JSON file that looks like:
//		{
//			"some": {
//				"value": 1
//			}
//		}
// The value can be accessed by calling the Get method:
//		val := cfg.GetInt("some.value", 1)
package json

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"strings"
)

type stringMap map[string]interface{}

// Configuration contains the loaded configuration and functions to manipulate it.
type Configuration struct {
	values stringMap
}

// NewConfiguration returns a new Configuration.
func NewConfiguration() *Configuration {
	return &Configuration{
		values: make(stringMap),
	}
}

// Load parses the JSON in data to it's internal map.
func (cfg *Configuration) Load(data string) error {
	m := make(stringMap)
	if err := json.Unmarshal([]byte(data), &m); err != nil {
		return err
	}
	cfg.values = cfg.loadMap(m)
	return nil
}

// loadMap parses a loaded map structure and adds it to the current configuration.
func (cfg Configuration) loadMap(m map[string]interface{}) stringMap {
	res := make(stringMap)
	for k, v := range m {
		if reflect.TypeOf(v).Kind() == reflect.Map {
			res[k] = cfg.loadMap(v.(map[string]interface{}))
		} else {
			res[k] = v
		}
	}
	return res
}

// Save returns the current configuration in JSON format.
func (cfg Configuration) Save() (string, error) {
	b, err := json.Marshal(&cfg.values)
	return string(b), err
}

// LoadFile reads a specified file into memory and parses it using Load().
func (cfg *Configuration) LoadFile(path string) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	jsonData := string(b[:])
	return cfg.Load(jsonData)
}

// SaveFile writes the output from Save() to the specified file.
func (cfg Configuration) SaveFile(path string) error {
	yml, err := cfg.Save()
	if err == nil {
		err = ioutil.WriteFile(path, []byte(yml), 0644)
	}
	return err
}

// Set sets a given value in the configuration to an arbitrary type.
func (cfg Configuration) Set(key string, val interface{}) {
	cfg.values[key] = val
}

// Remove deletes the value with the given key from the configuration.
func (cfg Configuration) Remove(key string) {
	delete(cfg.values, key)
}

// Get returns the value of the given key from the current configuration or the value of def if not found.
// The value is returned as an interface{}.
func (cfg Configuration) Get(key string, def interface{}) interface{} {
	return cfg.getFromMap(cfg.values, key, def)
}

// getFromMap returns the value of the given key from the specified map or the value of def if not found.
// The value is returned as an interface{}.
// If the key contains a dot (.), the function will be called recursively.
func (cfg Configuration) getFromMap(m stringMap, key string, def interface{}) interface{} {
	if n := strings.Index(key, "."); n != -1 {
		return cfg.getFromMap(m[key[:n]].(stringMap), key[n+1:], def)
	}
	if v, found := m[key]; found {
		return v
	}
	return def
}

// GetInt returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as an int.
func (cfg Configuration) GetInt(key string, def int) int {
	return cfg.Get(key, def).(int)
}

// GetInt8 returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as an int8.
func (cfg Configuration) GetInt8(key string, def int8) int8 {
	return int8(cfg.Get(key, def).(int))
}

// GetInt16 returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as an int16.
func (cfg Configuration) GetInt16(key string, def int16) int16 {
	return int16(cfg.Get(key, def).(int))
}

// GetInt32 returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as an int32.
func (cfg Configuration) GetInt32(key string, def int32) int32 {
	return int32(cfg.Get(key, def).(int))
}

// GetInt64 returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as an int64.
func (cfg Configuration) GetInt64(key string, def int64) int64 {
	return int64(cfg.Get(key, def).(int))
}

// GetUint returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as an uint.
func (cfg Configuration) GetUint(key string, def uint) uint {
	return uint(cfg.Get(key, def).(int))
}

// GetUint8 returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as an uint8.
func (cfg Configuration) GetUint8(key string, def uint8) uint8 {
	return uint8(cfg.Get(key, def).(int))
}

// GetUint16 returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as an uint16.
func (cfg Configuration) GetUint16(key string, def uint16) uint16 {
	return uint16(cfg.Get(key, def).(int))
}

// GetUint32 returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as an uint32.
func (cfg Configuration) GetUint32(key string, def uint32) uint32 {
	return uint32(cfg.Get(key, def).(int))
}

// GetUint64 returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as an uint64.
func (cfg Configuration) GetUint64(key string, def uint64) uint64 {
	return uint64(cfg.Get(key, def).(int))
}

// GetFloat32 returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as a float32.
func (cfg Configuration) GetFloat32(key string, def float32) float32 {
	return float32(cfg.Get(key, def).(float64))
}

// GetFloat64 returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as a float64.
func (cfg Configuration) GetFloat64(key string, def float64) float64 {
	return cfg.Get(key, def).(float64)
}

// GetBool returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as a bool.
func (cfg Configuration) GetBool(key string, def bool) bool {
	return cfg.Get(key, def).(bool)
}

// GetString returns the value of the given key from the current configuration or the value of def if not found.
// The value is returns as a string.
func (cfg Configuration) GetString(key string, def string) string {
	return cfg.Get(key, def).(string)
}
