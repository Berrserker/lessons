// Usage:
// func() {
// GetBuffer()
// defer PutBuffer()
// }

package buffer

import (
	"bytes"
	"sync"
)

var bufferPool sync.Pool

func GetBuffer() *bytes.Buffer {
	b := bufferPool.Get()
	if b != nil {
		// vb := b.(*bytes.Buffer)
		// vb.Reset()
		// return vb
		return b.(*bytes.Buffer)
	}
	return &bytes.Buffer{}
}

func PutBuffer(b *bytes.Buffer ) {

	if b == nil {
		return
	}
	b.Reset()
	bufferPool.Put(b)
}


