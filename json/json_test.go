package json

import (
	"io/ioutil"
	"os"
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
