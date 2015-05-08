package json

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type JSONSuite struct {
	data string
	file *os.File
}

var _ = Suite(&JSONSuite{})

func (s *JSONSuite) SetUpTest(c *C) {
	s.file, _ = ioutil.TempFile("", "config_test")
	s.file.Close()
	vars["num"] = 12345
	vars["f"] = 1.2345
	vars["str"] = "some string"
	vars["b"] = true
	vars["nested"] = make(stringMap)
	vars["nested"].(stringMap)["value"] = 1
}

func (s *JSONSuite) TearDownTest(c *C) {
	os.Remove(s.file.Name())
}

func (s *JSONSuite) TestLoad(c *C) {
	json := `{"test": "value"}`
	err := Load(json)
	_, found := vars["test"]
	c.Check(err, Equals, nil)
	c.Check(found, Equals, true)
	json = `{illegaljson, "test": "value"}`
	err = Load(json)
	c.Check(err, Not(Equals), nil)
}

func (s *JSONSuite) TestSave(c *C) {
	json, err := Save()
	c.Check(err, Equals, nil)
	c.Check(len(json), Not(Equals), 0)
}

func (s *JSONSuite) TestLoadFile(c *C) {
	json := `{"test": "value", "other": {"nested": 0}}`
	f, _ := os.Create(s.file.Name())
	f.WriteString(json)
	f.Close()
	err := LoadFile(s.file.Name())
	_, found := vars["test"]
	c.Check(err, Equals, nil)
	c.Check(found, Equals, true)
	err = LoadFile("nonexistant.file")
	c.Check(err, Not(Equals), nil)
}

func (s *JSONSuite) TestSaveFile(c *C) {
	err := SaveFile(s.file.Name())
	fi, _ := os.Stat(s.file.Name())
	c.Check(err, Equals, nil)
	c.Check(fi.Size(), Not(Equals), 0)
}

func (s *JSONSuite) TestSet(c *C) {
	Set("test", "value")
	v, found := vars["test"]
	c.Check(found, Equals, true)
	c.Check(v, Equals, "value")
}

func (s *JSONSuite) TestRemove(c *C) {
	vars["test"] = "value"
	Remove("test")
	_, found := vars["test"]
	c.Check(found, Equals, false)
}

func (s *JSONSuite) TestGet(c *C) {
	vars["test"] = "value"
	v := Get("test", "default")
	c.Check(v, Equals, "value")
	v = Get("nested.value", "default")
	c.Check(v, Not(Equals), "default")
	v = Get("nonexistant", "default")
	c.Check(v, Equals, "default")
	v = Get("nonexistant.nested", "default")
	c.Check(v, Equals, "default")
}

func (s *JSONSuite) TestGetInt(c *C) {
	v := GetInt("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Int)
	c.Check(v, Not(Equals), 0)
}

func (s *JSONSuite) TestGetInt8(c *C) {
	v := GetInt8("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Int8)
	c.Check(v, Not(Equals), 0)
}

func (s *JSONSuite) TestGetInt16(c *C) {
	v := GetInt16("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Int16)
	c.Check(v, Not(Equals), 0)
}

func (s *JSONSuite) TestGetInt32(c *C) {
	v := GetInt32("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Int32)
	c.Check(v, Not(Equals), 0)
}

func (s *JSONSuite) TestGetInt64(c *C) {
	v := GetInt64("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Int64)
	c.Check(v, Not(Equals), 0)
}

func (s *JSONSuite) TestGetUint(c *C) {
	v := GetUint("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Uint)
	c.Check(v, Not(Equals), 0)
}

func (s *JSONSuite) TestGetUint8(c *C) {
	v := GetUint8("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Uint8)
	c.Check(v, Not(Equals), 0)
}

func (s *JSONSuite) TestGetUint16(c *C) {
	v := GetUint16("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Uint16)
	c.Check(v, Not(Equals), 0)
}

func (s *JSONSuite) TestGetUint32(c *C) {
	v := GetUint32("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Uint32)
	c.Check(v, Not(Equals), 0)
}

func (s *JSONSuite) TestGetUint64(c *C) {
	v := GetUint64("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Uint64)
	c.Check(v, Not(Equals), 0)
}

func (s *JSONSuite) TestGetFloat32(c *C) {
	v := GetFloat32("f", 0.0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Float32)
	c.Check(v, Not(Equals), 0.0)
}

func (s *JSONSuite) TestGetFloat64(c *C) {
	v := GetFloat64("f", 0.0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Float64)
	c.Check(v, Not(Equals), 0.0)
}

func (s *JSONSuite) TestGetBool(c *C) {
	v := GetBool("b", false)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Bool)
	c.Check(v, Not(Equals), false)
}

func (s *JSONSuite) TestGetString(c *C) {
	v := GetString("str", "default")
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.String)
	c.Check(v, Not(Equals), "default")
}

func (s *JSONSuite) BenchmarkConfigurationLoad(c *C) {
	for i := 0; i < c.N; i++ {
		Load(s.data)
	}
}

func (s *JSONSuite) BenchmarkConfigurationSave(c *C) {
	for i := 0; i < c.N; i++ {
		Save()
	}
}

func (s *JSONSuite) BenchmarkConfigurationGet(c *C) {
	for i := 0; i < c.N; i++ {
		Get("str", "default")
	}
}

func (s *JSONSuite) BenchmarkConfigurationGetNested(c *C) {
	for i := 0; i < c.N; i++ {
		Get("nested.value", "default")
	}
}
