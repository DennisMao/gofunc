package json

import (
	j "encoding/json"
	"testing"
)

type Raw struct {
	Id   int64       `json:"raw"`
	Name string      `json:"name"`
	Data interface{} `json:"data"`
	Cell float64     `json:"cell"`
}

func BenchmarkDidiJsonMarshal(b *testing.B) {
	r := Raw{0, "1234567890", "1234567890", 0.0}
	for i := 0; i < b.N; i++ {
		Marshal(r)
	}
}

func BenchmarkOfficialJsonMarshal(b *testing.B) {
	r := Raw{0, "1234567890", "1234567890", 0.0}
	for i := 0; i < b.N; i++ {
		j.Marshal(r)
	}
}

func BenchmarkDidiJsonUnmarshal(b *testing.B) {
	text := []byte(`{"id":0,"name":"1234567890","data":"1234567890","cell":0.0}`)
	r := Raw{}
	for i := 0; i < b.N; i++ {
		Unmarshal(text, &r)
	}
}

func BenchmarkOfficialJsonUnmarshal(b *testing.B) {
	text := []byte(`{"id":0,"name":"1234567890","data":"1234567890","cell":0.0}`)
	r := Raw{}
	for i := 0; i < b.N; i++ {
		j.Unmarshal(text, &r)
	}
}
