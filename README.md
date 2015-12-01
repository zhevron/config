config - Simple configuration file library
==========================================

[![wercker status](https://app.wercker.com/status/e7ff76576fa4b805c007f48a2377a925/s/master "wercker status")](https://app.wercker.com/project/bykey/e7ff76576fa4b805c007f48a2377a925)
[![Coverage Status](https://coveralls.io/repos/zhevron/config/badge.svg?branch=master&service=github)](https://coveralls.io/github/zhevron/config?branch=master)
[![GoDoc](https://godoc.org/github.com/zhevron/config?status.svg)](https://godoc.org/github.com/zhevron/config)

**config** is a simple configuration file library for [Google Go](https://golang.org).  
For full package documentation, see the GoDoc link above.

## Usage

### Get/Set values
```go
config.GetString("var", "default")
config.Set("var", value)
```

### Load/Save JSON
```go
if err := config.LoadFileJSON(filename); err != nil {
  panic(err) // Failed to load configuration file
}
if err := config.SaveFileJSON(filename); err != nil {
  panic(err) // Failed to save configuration file
}
```

### Load/Save YAML
```go
if err := config.LoadFileYAML(filename); err != nil {
  panic(err) // Failed to load configuration file
}
if err := config.SaveFileYAML(filename); err != nil {
  panic(err) // Failed to save configuration file
}
```

## License

**config** is licensed under the [MIT license](http://opensource.org/licenses/MIT).
