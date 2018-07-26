package int128

import (
	"math/big"
	//	"math/rand"
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	var x = 9223372036854543385
	var y = 9223372036854775574
	var z int

	b.Run("Simple", func(b *testing.B) {
		for i := 0; i != b.N; i++ {
			z = x + y
		}
	})
}

func BenchmarkBigAdd(b *testing.B) {
	var x = big.NewInt(9223372036854543385)
	var y = big.NewInt(9223372036854775574)
	var z big.Int

	b.Run("Simple", func(b *testing.B) {
		for i := 0; i != b.N; i++ {
			z.Add(x, y)
		}
	})
}

func BenchmarkInt128Add(b *testing.B) {
	var x = Uint128{2, 9223372036854543385}
	var y = Uint128{3, 9223372036854775574}
	var z Uint128

	b.Run("Simple", func(b *testing.B) {
		for i := 0; i != b.N; i++ {
			z = Add(x, y)
		}
	})
}
