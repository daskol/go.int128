package int128

import (
	"testing"
)

func TestUint128Add(t *testing.T) {
	var x = Uint128{7, 1}
	var y = Uint128{4, 2}

	t.Run("Correctness", func(t *testing.T) {
		var z = Add(x, y)

		if z.hi != x.hi+y.hi {
			t.Errorf("wrong hi dword: %d != %d + %d", z.hi, x.hi, y.hi)
		}

		if z.lo != x.lo+y.lo {
			t.Errorf("wrong lo dword: %d != %d + %d", z.lo, x.lo, y.lo)
		}
	})
}
