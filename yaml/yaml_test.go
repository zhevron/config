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

type GocfgYamlSuite struct {
	cfg *Configuration
	yml string
	f   *os.File
}

var _ = Suite(&GocfgYamlSuite{})

func (s *GocfgYamlSuite) SetUpTest(c *C) {
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

func (s *GocfgYamlSuite) TearDownTest(c *C) {
	os.Remove(s.f.Name())
}

func (s *GocfgYamlSuite) TestNewConfiguration(c *C) {
	cfg := NewConfiguration()
	c.Check(cfg, Not(Equals), nil)
}

func (s *GocfgYamlSuite) TestConfigurationLoad(c *C) {
	yml := "test: value"
	err := s.cfg.Load(yml)
	_, found := s.cfg.values["test"]
	c.Check(err, Equals, nil)
	c.Check(found, Equals, true)
	yml = `
illegalyaml
test: value
`
	err = s.cfg.Load(yml)
	c.Check(err, Not(Equals), nil)
}

func (s *GocfgYamlSuite) TestConfigurationSave(c *C) {
	yml, err := s.cfg.Save()
	c.Check(err, Equals, nil)
	c.Check(len(yml), Not(Equals), 0)
}

func (s *GocfgYamlSuite) TestConfigurationLoadFile(c *C) {
	yml := `
test: value
other:
  nested: 0
`
	f, _ := os.Create(s.f.Name())
	f.WriteString(yml)
	f.Close()
	err := s.cfg.LoadFile(s.f.Name())
	_, found := s.cfg.values["test"]
	c.Check(err, Equals, nil)
	c.Check(found, Equals, true)
	err = s.cfg.LoadFile("nonexistant.file")
	c.Check(err, Not(Equals), nil)
}

func (s *GocfgYamlSuite) TestConfigurationSaveFile(c *C) {
	err := s.cfg.SaveFile(s.f.Name())
	fi, _ := os.Stat(s.f.Name())
	c.Check(err, Equals, nil)
	c.Check(fi.Size(), Not(Equals), 0)
}

func (s *GocfgYamlSuite) TestConfigurationSet(c *C) {
	s.cfg.Set("test", "value")
	v, found := s.cfg.values["test"]
	c.Check(found, Equals, true)
	c.Check(v, Equals, "value")
}

func (s *GocfgYamlSuite) TestConfigurationRemove(c *C) {
	s.cfg.values["test"] = "value"
	s.cfg.Remove("test")
	_, found := s.cfg.values["test"]
	c.Check(found, Equals, false)
}

func (s *GocfgYamlSuite) TestConfigurationGet(c *C) {
	s.cfg.values["test"] = "value"
	v := s.cfg.Get("test", "default")
	c.Check(v, Equals, "value")
	v = s.cfg.Get("nested.value", "default")
	c.Check(v, Not(Equals), "default")
	v = s.cfg.Get("nonexistant", "default")
	c.Check(v, Equals, "default")
}

func (s *GocfgYamlSuite) TestConfigurationGetInt(c *C) {
	v := s.cfg.GetInt("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Int)
	c.Check(v, Not(Equals), 0)
}

func (s *GocfgYamlSuite) TestConfigurationGetInt8(c *C) {
	v := s.cfg.GetInt8("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Int8)
	c.Check(v, Not(Equals), 0)
}

func (s *GocfgYamlSuite) TestConfigurationGetInt16(c *C) {
	v := s.cfg.GetInt16("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Int16)
	c.Check(v, Not(Equals), 0)
}

func (s *GocfgYamlSuite) TestConfigurationGetInt32(c *C) {
	v := s.cfg.GetInt32("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Int32)
	c.Check(v, Not(Equals), 0)
}

func (s *GocfgYamlSuite) TestConfigurationGetInt64(c *C) {
	v := s.cfg.GetInt64("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Int64)
	c.Check(v, Not(Equals), 0)
}

func (s *GocfgYamlSuite) TestConfigurationGetUint(c *C) {
	v := s.cfg.GetUint("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Uint)
	c.Check(v, Not(Equals), 0)
}

func (s *GocfgYamlSuite) TestConfigurationGetUint8(c *C) {
	v := s.cfg.GetUint8("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Uint8)
	c.Check(v, Not(Equals), 0)
}

func (s *GocfgYamlSuite) TestConfigurationGetUint16(c *C) {
	v := s.cfg.GetUint16("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Uint16)
	c.Check(v, Not(Equals), 0)
}

func (s *GocfgYamlSuite) TestConfigurationGetUint32(c *C) {
	v := s.cfg.GetUint32("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Uint32)
	c.Check(v, Not(Equals), 0)
}

func (s *GocfgYamlSuite) TestConfigurationGetUint64(c *C) {
	v := s.cfg.GetUint64("num", 0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Uint64)
	c.Check(v, Not(Equals), 0)
}

func (s *GocfgYamlSuite) TestConfigurationGetFloat32(c *C) {
	v := s.cfg.GetFloat32("f", 0.0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Float32)
	c.Check(v, Not(Equals), 0.0)
}

func (s *GocfgYamlSuite) TestConfigurationGetFloat64(c *C) {
	v := s.cfg.GetFloat64("f", 0.0)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Float64)
	c.Check(v, Not(Equals), 0.0)
}

func (s *GocfgYamlSuite) TestConfigurationGetBool(c *C) {
	v := s.cfg.GetBool("b", false)
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.Bool)
	c.Check(v, Not(Equals), false)
}

func (s *GocfgYamlSuite) TestConfigurationGetString(c *C) {
	v := s.cfg.GetString("str", "default")
	c.Check(reflect.TypeOf(v).Kind(), Equals, reflect.String)
	c.Check(v, Not(Equals), "default")
}

func (s *GocfgYamlSuite) BenchmarkConfigurationLoad(c *C) {
	for i := 0; i < c.N; i++ {
		s.cfg.Load(s.yml)
	}
}

func (s *GocfgYamlSuite) BenchmarkConfigurationSave(c *C) {
	for i := 0; i < c.N; i++ {
		s.cfg.Save()
	}
}

func (s *GocfgYamlSuite) BenchmarkConfigurationGet(c *C) {
	for i := 0; i < c.N; i++ {
		s.cfg.Get("str", "default")
	}
}

func (s *GocfgYamlSuite) BenchmarkConfigurationGetNested(c *C) {
	for i := 0; i < c.N; i++ {
		s.cfg.Get("nested.value", "default")
	}
}
