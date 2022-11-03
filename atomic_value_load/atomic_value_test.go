package atomic_value_load

import (
	"github.com/stretchr/testify/assert"
	"sync/atomic"
	"testing"
)

var sintp = new(int)
var atomicVal = atomic.Value{}

func TestNotNull(t *testing.T) {
	assert.NotNil(t, sintp)
}

func BenchmarkAtomicLoad(b *testing.B) {
	b.Run("AccessGlobalValue", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = sintp
			}
		})
	})

	b.Run("AccessGlobalValue", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		atomicVal.Store(sintp)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = atomicVal.Load().(*int)
			}
		})
	})
}
