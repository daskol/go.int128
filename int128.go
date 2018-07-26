package int128

type Uint128 struct {
	hi uint64
	lo uint64
}

//go:noescape
//go:nosplit

func Add(x, y Uint128) (z Uint128)
