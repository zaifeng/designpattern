package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInstance(t *testing.T) {
	assert.Equal(t, GetInstance(), GetInsOnce())
	assert.Equal(t, GetInstance(), GetInsLazy())
	assert.Equal(t, GetInsLazy(), GetInsOnce())
}

func BenchmarkGetInstance(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			if GetInsLazy() != GetInsOnce() {
				b.Error("GetInstance benchmark test failed")
			}

			if GetInstance() != GetInsOnce() {
				b.Error("GetInstance benchmark test failed")
			}

			if GetInstance() != GetInsLazy() {
				b.Error("GetInstance benchmark test failed")
			}
		}
	})
}
