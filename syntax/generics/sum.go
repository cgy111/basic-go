package main

import (
	"encoding/json"
)

func Sum[T Num](vals ...T) T {
	var res T
	for _, val := range vals {
		res = res + val
	}
	return res
}

type Num interface {
	~int | int64 | float64
}

type Integer int

func ReleaseResource[R json.Marshaler](r R) {
	r.MarshalJSON()
}
