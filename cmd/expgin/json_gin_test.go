package expgin

import (
	"gopkg.in/h2non/baloo.v3"
	"testing"
)

const schema = `{
  "title": "Example Schema",
  "type": "object",
  "properties": {
    "origin": {
      "type": "string"
    }
  },
  "required": ["origin"]
}`

//const port = ":10001"
//var testjson = baloo.New("http://127.0.0.1"+port)

var testbin = baloo.New("http://httpbin.org")

func TestJsonBin(t *testing.T) {
	testbin.Get("/ip").
		Expect(t).
		Status(200).
		Type("json").
		JSONSchema(schema).
		Done()

	t.Log(string(schema))
}

