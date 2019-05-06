package statistics

import (
	"encoding/json"
)

type Kvs map[string]string
var allkvs = make(map[int64]Kvs)
var token string

var groups []int64


func jsonify(kvsin Kvs) string {
	s, err :=json.Marshal(kvsin)
	checkErr(err)
	return string(s)
}

func json2kvs(jsonin string) Kvs {
	tkvs := make(Kvs)
	_ = json.Unmarshal([]byte(jsonin), &tkvs)
	return tkvs
}


func loadinfo() {
	dbread()
}
