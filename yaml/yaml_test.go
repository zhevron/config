package yaml

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

type YAMLSuite struct {
	data string
	file *os.File
}

var _ = Suite(&YAMLSuite{})

func (s *YAMLSuite) SetUpTest(c *C) {
	s.file, _ = ioutil.TempFile("", "gocfg")
	s.file.Close()
	vars["num"] = 12345
	vars["f"] = 1.2345
	vars["str"] = "some string"
	vars["b"] = true
	vars["nested"] = make(stringMap)
	vars["nested"].(stringMap)["value"] = 1
}

func (s *YAMLSuite) TearDownTest(c *C) {
	os.Remove(s.file.Name())
}

func (s *YAMLSuite) TestLoad(c *C) {
	yml := "test: value"
	err := Load(yml)
	_, found := vars["test"]
	c.Check(err, Equals, nil)
	c.Check(found, Equals, true)
	yml = `
illegalyaml
test: value
`
	err = Load(yml)
	c.Check(err, Not(Equals), nil)
}

func (s *YAMLSuite) TestSave(c *C) {
	yml, err := Save()
	c.Check(err, Equals, nil)
	c.Check(len(yml), Not(Equals), 0)
}

func (s *YAMLSuite) TestLoadFile(c *C) {
	yml := `
test: value
other:
  nested: 0
`
	f, _ := os.Create(s.file.Name())
	f.WriteString(yml)
	f.Close()
	err := LoadFile(s.file.Name())
	_, found := vars["test"]
	c.Check(err, Equals, nil)
	c.Check(found, Equals, true)
	err = LoadFile("nonexistant.file")
	c.Check(err, Not(Equals), nil)
}

func (s *YAMLSuite) TestSaveFile(c *C) {
	err := SaveFile(s.file.Name())
	fi, _ := os.Stat(s.file.Name())
	c.Check(err, Equals, nil)
	c.Check(fi.Size(), Not(Equals), 0)
}

func (s *YAMLSuite) TestSet(c *C) {
	Set("test", "value")
	v, found := vars["test"]
	c.Check(found, Equals, true)
	c.Check(v, Equals, "value")
}

func (s *YAMLSuite) TestRemove(c *C) {
	vars["test"] = "value"
	Remove("test")
	_, found := vars["test"]
	c.Check(found, Equals, false)
}

func (s *YAMLSuite) TestGet(c *C) {
	vars["test"] = "value"
	v := Get("test", "default")
	c.Check(v, Equals, "value")
	v = Get("nested.value", "default")
	c.Check(v, Not(Equals), "default")
	v = Get("nonexistant", "default")
	c.Check(v, Equals, "default")
}

func (s *YAMLSuite) TestGetInt(c *C) {
	v := GetInt("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Int)
	c.Check(v, Not(Equals), 0)
}

func (s *YAMLSuite) TestGetInt8(c *C) {
	v := GetInt8("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Int8)
	c.Check(v, Not(Equals), 0)
}

func (s *YAMLSuite) TestGetInt16(c *C) {
	v := GetInt16("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Int16)
	c.Check(v, Not(Equals), 0)
}

func (s *YAMLSuite) TestGetInt32(c *C) {
	v := GetInt32("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Int32)
	c.Check(v, Not(Equals), 0)
}

func (s *YAMLSuite) TestGetInt64(c *C) {
	v := GetInt64("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Int64)
	c.Check(v, Not(Equals), 0)
}

func (s *YAMLSuite) TestGetUint(c *C) {
	v := GetUint("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Uint)
	c.Check(v, Not(Equals), 0)
}

func (s *YAMLSuite) TestGetUint8(c *C) {
	v := GetUint8("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Uint8)
	c.Check(v, Not(Equals), 0)
}

func (s *YAMLSuite) TestGetUint16(c *C) {
	v := GetUint16("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Uint16)
	c.Check(v, Not(Equals), 0)
}

func (s *YAMLSuite) TestGetUint32(c *C) {
	v := GetUint32("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Uint32)
	c.Check(v, Not(Equals), 0)
}

func (s *YAMLSuite) TestGetUint64(c *C) {
	v := GetUint64("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Uint64)
	c.Check(v, Not(Equals), 0)
}

func (s *YAMLSuite) TestGetFloat32(c *C) {
	v := GetFloat32("f", 0.0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Float32)
	c.Check(v, Not(Equals), 0.0)
}

func (s *YAMLSuite) TestGetFloat64(c *C) {
	v := GetFloat64("f", 0.0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Float64)
	c.Check(v, Not(Equals), 0.0)
}

func (s *YAMLSuite) TestGetBool(c *C) {
	v := GetBool("b", false)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Bool)
	c.Check(v, Not(Equals), false)
}

func (s *YAMLSuite) TestGetString(c *C) {
	v := GetString("str", "default")
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.String)
	c.Check(v, Not(Equals), "default")
}

func (s *YAMLSuite) BenchmarkConfigurationLoad(c *C) {
	for i := 0; i < c.N; i++ {
		Load(s.data)
	}
}

func (s *YAMLSuite) BenchmarkConfigurationSave(c *C) {
	for i := 0; i < c.N; i++ {
		Save()
	}
}

func (s *YAMLSuite) BenchmarkConfigurationGet(c *C) {
	for i := 0; i < c.N; i++ {
		Get("str", "default")
	}
}

func (s *YAMLSuite) BenchmarkConfigurationGetNested(c *C) {
	for i := 0; i < c.N; i++ {
		Get("nested.value", "default")
	}
}
