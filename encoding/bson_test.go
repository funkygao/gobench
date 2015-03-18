package encoding

import (
	"encoding/json"
	"labix.org/v2/mgo/bson"
	"testing"
)

func BenchmarkJson(b *testing.B) {
	b.ReportAllocs()
	m := bson.M{
		"userId": 343434,
		"gendar": "F",
		"info": bson.M{
			"city":    "beijing",
			"hobbies": []string{"a", "b"}}}
	var v []byte
	for i := 0; i < b.N; i++ {
		v, _ = json.Marshal(m)
	}
	b.SetBytes(int64(len(v)))
}

func BenchmarkBson(b *testing.B) {
	b.ReportAllocs()
	m := bson.M{
		"userId": 343434,
		"gendar": "F",
		"info": bson.M{
			"city":    "beijing",
			"hobbies": []string{"a", "b"}}}
	var v []byte
	for i := 0; i < b.N; i++ {
		v, _ = bson.Marshal(m)
	}
	b.SetBytes(int64(len(v)))
}
