// Copyright (C) 2014 Thomas Lokshall
// Use of this source code is governed by the MIT license.
// See LICENSE.md for details.

// Package json implements a way to easily read and write JSON configuration files.
package json

type interfaceMap map[interface{}]interface{}
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
