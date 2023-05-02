package common

import (
	"sync"
)

const (
	bufSize = 1024
)

var bpool sync.Pool

func init() {
	bpool.New = func() interface{} {
		return make([]byte, bufSize)
	}
}
func BufferPoolGet() []byte {
	return bpool.Get().([]byte)
}
func BufferPoolPut(b []byte) {
	bpool.Put(b)
}
