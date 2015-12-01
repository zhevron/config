package config

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/zhevron/match"
)

var file *os.File

func TestMain(m *testing.M) {
	file, _ = ioutil.TempFile("", "config_test")
	file.Close()
	defer os.Remove(file.Name())
	vars["num"] = 12345
	vars["f"] = 1.2345
	vars["str"] = "some string"
	vars["b"] = true
	vars["nested"] = make(stringMap)
	vars["nested"].(stringMap)["value"] = 1
	os.Exit(m.Run())
}

func TestSet(t *testing.T) {
	Set("test", "value")
	v, found := vars["test"]
	match.Equals(t, found, true)
	match.Equals(t, v, "value")
}

func TestRemove(t *testing.T) {
	vars["test"] = "value"
	Remove("test")
	_, found := vars["test"]
	match.Equals(t, found, false)
}

func TestGet(t *testing.T) {
	vars["test"] = "value"
	v := Get("test", "default")
	match.Equals(t, v, "value")
	v = Get("nested.value", "default")
	match.NotEquals(t, v, "default")
	v = Get("nonexistant", "default")
	match.Equals(t, v, "default")
	v = Get("nonexistant.nested", "default")
	match.Equals(t, v, "default")
	v = Get("test.nested", "default")
	match.Equals(t, v, "default")
}

func TestGet_Int(t *testing.T) {
	v := GetInt("num", 0)
	match.Equals(t, reflect.TypeOf(v).Kind(), reflect.Int)
	match.NotEquals(t, v, 0)
}

func TestGet_Int8(t *testing.T) {
	v := GetInt8("num", 0)
	match.Equals(t, reflect.TypeOf(v).Kind(), reflect.Int8)
	match.NotEquals(t, v, 0)
}

func TestGet_Int16(t *testing.T) {
	v := GetInt16("num", 0)
	match.Equals(t, reflect.TypeOf(v).Kind(), reflect.Int16)
	match.NotEquals(t, v, 0)
}

func TestGet_Int32(t *testing.T) {
	v := GetInt32("num", 0)
	match.Equals(t, reflect.TypeOf(v).Kind(), reflect.Int32)
	match.NotEquals(t, v, 0)
}

func TestGet_Int64(t *testing.T) {
	v := GetInt64("num", 0)
	match.Equals(t, reflect.TypeOf(v).Kind(), reflect.Int64)
	match.NotEquals(t, v, 0)
}

func TestGet_Uint(t *testing.T) {
	v := GetUint("num", 0)
	match.Equals(t, reflect.TypeOf(v).Kind(), reflect.Uint)
	match.NotEquals(t, v, 0)
}

func TestGet_Uint8(t *testing.T) {
	v := GetUint8("num", 0)
	match.Equals(t, reflect.TypeOf(v).Kind(), reflect.Uint8)
	match.NotEquals(t, v, 0)
}

func TestGet_Uint16(t *testing.T) {
	v := GetUint16("num", 0)
	match.Equals(t, reflect.TypeOf(v).Kind(), reflect.Uint16)
	match.NotEquals(t, v, 0)
}

func TestGet_Uint32(t *testing.T) {
	v := GetUint32("num", 0)
	match.Equals(t, reflect.TypeOf(v).Kind(), reflect.Uint32)
	match.NotEquals(t, v, 0)
}

func TestGet_Uint64(t *testing.T) {
	v := GetUint64("num", 0)
	match.Equals(t, reflect.TypeOf(v).Kind(), reflect.Uint64)
	match.NotEquals(t, v, 0)
}

func TestGet_Float32(t *testing.T) {
	v := GetFloat32("f", 0.0)
	match.Equals(t, reflect.TypeOf(v).Kind(), reflect.Float32)
	match.NotEquals(t, v, 0.0)
}

func TestGet_Float64(t *testing.T) {
	v := GetFloat64("f", 0.0)
	match.Equals(t, reflect.TypeOf(v).Kind(), reflect.Float64)
	match.NotEquals(t, v, 0.0)
}

func TestGet_Bool(t *testing.T) {
	v := GetBool("b", false)
	match.Equals(t, reflect.TypeOf(v).Kind(), reflect.Bool)
	match.NotEquals(t, v, false)
}

func TestGet_String(t *testing.T) {
	v := GetString("str", "default")
	match.Equals(t, reflect.TypeOf(v).Kind(), reflect.String)
	match.NotEquals(t, v, "default")
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get("str", "default")
	}
}

func BenchmarkGet_Nested(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get("nested.value", "default")
	}
}
