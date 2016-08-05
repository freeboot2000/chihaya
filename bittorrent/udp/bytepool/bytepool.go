// Copyright 2016 The Chihaya Authors. All rights reserved.
// Use of this source code is governed by the BSD 2-Clause license,
// which can be found in the LICENSE file.

package bytepool

import "sync"

// BytePool is a cached pool of reusable byte slices.
type BytePool struct {
	sync.Pool
}

// New allocates a new BytePool with slices of the provided capacity.
func New(length, capacity int) *BytePool {
	var bp BytePool
	bp.Pool.New = func() interface{} {
		return make([]byte, length, capacity)
	}
	return &bp
}

// Get returns a byte slice from the pool.
func (bp *BytePool) Get() []byte {
	return bp.Pool.Get().([]byte)
}

// Put returns a byte slice to the pool.
func (bp *BytePool) Put(b []byte) {
	// Zero out the bytes.
	for i := 0; i < cap(b); i++ {
		b[i] = 0x0
	}
	bp.Pool.Put(b)
}
