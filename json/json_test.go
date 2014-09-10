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

type GocfgJSONSuite struct {
	cfg *Configuration
	yml string
	f   *os.File
}

var _ = Suite(&GocfgJSONSuite{})

func (s *GocfgJSONSuite) SetUpTest(c *C) {
	s.cfg = NewConfiguration()
	s.f, _ = ioutil.TempFile("", "gocfg")
	s.f.Close()
	s.cfg.values["num"] = 12345
	s.cfg.values["f"] = 1.2345
	s.cfg.values["str"] = "some string"
	s.cfg.values["b"] = true
	s.cfg.values["nested"] = make(stringMap)
	s.cfg.values["nested"].(stringMap)["value"] = 1
}

func (s *GocfgJSONSuite) TearDownTest(c *C) {
	os.Remove(s.f.Name())
}

func (s *GocfgJSONSuite) TestNewConfiguration(c *C) {
	cfg := NewConfiguration()
	c.Check(cfg, Not(Equals), nil)
}

func (s *GocfgJSONSuite) TestConfigurationLoad(c *C) {
	json := `{"test": "value"}`
	err := s.cfg.Load(json)
	_, found := s.cfg.values["test"]
	c.Check(err, Equals, nil)
	c.Check(found, Equals, true)
	json = `{illegaljson, "test": "value"}`
	err = s.cfg.Load(json)
	c.Check(err, Not(Equals), nil)
}

func (s *GocfgJSONSuite) TestConfigurationSave(c *C) {
	json, err := s.cfg.Save()
	c.Check(err, Equals, nil)
	c.Check(len(json), Not(Equals), 0)
}

func (s *GocfgJSONSuite) TestConfigurationLoadFile(c *C) {
	json := `{"test": "value", "other": {"nested": 0}}`
	f, _ := os.Create(s.f.Name())
	f.WriteString(json)
	f.Close()
	err := s.cfg.LoadFile(s.f.Name())
	_, found := s.cfg.values["test"]
	c.Check(err, Equals, nil)
	c.Check(found, Equals, true)
	err = s.cfg.LoadFile("nonexistant.file")
	c.Check(err, Not(Equals), nil)
}

func (s *GocfgJSONSuite) TestConfigurationSaveFile(c *C) {
	err := s.cfg.SaveFile(s.f.Name())
	fi, _ := os.Stat(s.f.Name())
	c.Check(err, Equals, nil)
	c.Check(fi.Size(), Not(Equals), 0)
}

func (s *GocfgJSONSuite) TestConfigurationSet(c *C) {
	s.cfg.Set("test", "value")
	v, found := s.cfg.values["test"]
	c.Check(found, Equals, true)
	c.Check(v, Equals, "value")
}

func (s *GocfgJSONSuite) TestConfigurationRemove(c *C) {
	s.cfg.values["test"] = "value"
	s.cfg.Remove("test")
	_, found := s.cfg.values["test"]
	c.Check(found, Equals, false)
}

func (s *GocfgJSONSuite) TestConfigurationGet(c *C) {
	s.cfg.values["test"] = "value"
	v := s.cfg.Get("test", "default")
	c.Check(v, Equals, "value")
	v = s.cfg.Get("nested.value", "default")
	c.Check(v, Not(Equals), "default")
	v = s.cfg.Get("nonexistant", "default")
	c.Check(v, Equals, "default")
}

func (s *GocfgJSONSuite) TestConfigurationGetInt(c *C) {
	v := s.cfg.GetInt("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Int)
	c.Check(v, Not(Equals), 0)
}

func (s *GocfgJSONSuite) TestConfigurationGetInt8(c *C) {
	v := s.cfg.GetInt8("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Int8)
	c.Check(v, Not(Equals), 0)
}

func (s *GocfgJSONSuite) TestConfigurationGetInt16(c *C) {
	v := s.cfg.GetInt16("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Int16)
	c.Check(v, Not(Equals), 0)
}

func (s *GocfgJSONSuite) TestConfigurationGetInt32(c *C) {
	v := s.cfg.GetInt32("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Int32)
	c.Check(v, Not(Equals), 0)
}

func (s *GocfgJSONSuite) TestConfigurationGetInt64(c *C) {
	v := s.cfg.GetInt64("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Int64)
	c.Check(v, Not(Equals), 0)
}

func (s *GocfgJSONSuite) TestConfigurationGetUint(c *C) {
	v := s.cfg.GetUint("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Uint)
	c.Check(v, Not(Equals), 0)
}

func (s *GocfgJSONSuite) TestConfigurationGetUint8(c *C) {
	v := s.cfg.GetUint8("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Uint8)
	c.Check(v, Not(Equals), 0)
}

func (s *GocfgJSONSuite) TestConfigurationGetUint16(c *C) {
	v := s.cfg.GetUint16("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Uint16)
	c.Check(v, Not(Equals), 0)
}

func (s *GocfgJSONSuite) TestConfigurationGetUint32(c *C) {
	v := s.cfg.GetUint32("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Uint32)
	c.Check(v, Not(Equals), 0)
}

func (s *GocfgJSONSuite) TestConfigurationGetUint64(c *C) {
	v := s.cfg.GetUint64("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Uint64)
	c.Check(v, Not(Equals), 0)
}

func (s *GocfgJSONSuite) TestConfigurationGetFloat32(c *C) {
	v := s.cfg.GetFloat32("f", 0.0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Float32)
	c.Check(v, Not(Equals), 0.0)
}

func (s *GocfgJSONSuite) TestConfigurationGetFloat64(c *C) {
	v := s.cfg.GetFloat64("f", 0.0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Float64)
	c.Check(v, Not(Equals), 0.0)
}

func (s *GocfgJSONSuite) TestConfigurationGetBool(c *C) {
	v := s.cfg.GetBool("b", false)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Bool)
	c.Check(v, Not(Equals), false)
}

func (s *GocfgJSONSuite) TestConfigurationGetString(c *C) {
	v := s.cfg.GetString("str", "default")
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.String)
	c.Check(v, Not(Equals), "default")
}

func (s *GocfgJSONSuite) BenchmarkConfigurationLoad(c *C) {
	for i := 0; i < c.N; i++ {
		s.cfg.Load(s.yml)
	}
}

func (s *GocfgJSONSuite) BenchmarkConfigurationSave(c *C) {
	for i := 0; i < c.N; i++ {
		s.cfg.Save()
	}
}

func (s *GocfgJSONSuite) BenchmarkConfigurationGet(c *C) {
	for i := 0; i < c.N; i++ {
		s.cfg.Get("str", "default")
	}
}

func (s *GocfgJSONSuite) BenchmarkConfigurationGetNested(c *C) {
	for i := 0; i < c.N; i++ {
		s.cfg.Get("nested.value", "default")
	}
}
