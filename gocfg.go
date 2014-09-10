// Copyright (C) 2014 Thomas Lokshall
// Use of this source code is governed by the MIT license.
// See LICENSE.md for details.

// Package gocfg implements a way to easily read and write several types of configuration files.
//
// This package should not be used. Instead, use one of the sub-packages.
package gocfg

// Configuration is the base interface that all sub-packages implement.
type Configuration interface {
	Load(string) error
	Save() (string, error)
	LoadFile(string) error
	SaveFile(string) error
	Set(string, interface{})
	Remove(string)
	Get(string, interface{}) interface{}
	GetInt(string, int) int
	GetInt8(string, int8) int8
	GetInt16(string, int16) int16
	GetInt32(string, int32) int32
	GetInt64(string, int64) int64
	GetUint(string, uint) uint
	GetUint8(string, uint8) uint8
	GetUint16(string, uint16) uint16
	GetUint32(string, uint32) uint32
	GetUint64(string, uint64) uint64
	GetFloat32(string, float32) float32
	GetFloat64(string, float64) float64
	GetBool(string, bool) bool
	GetString(string, string) string
}
