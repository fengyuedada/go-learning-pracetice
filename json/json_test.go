package json

import (
	"encoding/json"
	"testing"
)


var jsonStr = `{
	"basic_info":{
		"name":"Mike",
		"age":30
	},
	"job_info":{
		"skills":["Java","Go"]
	}
}`

func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(*e)
	v, err := json.Marshal(e)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(v))
}

func TestEasyJson(t *testing.T) {
	e := Employee{}
	e.UnmarshalJSON([]byte(jsonStr))
	t.Log(e)
	v, err := e.MarshalJSON()
	if err != nil {
		t.Log(err)
	}
	t.Log(string(v))
}

func BenchmarkEmbeddedJson(b *testing.B) {
	b.ResetTimer()
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal([]byte(jsonStr), e)
		if err != nil {
			b.Fatal(err)
		}
		_, err = json.Marshal(e)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEasyJson(b *testing.B) {
	b.ResetTimer()
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		e.UnmarshalJSON([]byte(jsonStr))
		_, err := e.MarshalJSON()
		if err != nil {
			b.Log(err)
		}
	}
}