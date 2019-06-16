package hashmap

import "github.com/syafdia/go-exercise/src/types"

type HashMap interface {
	Set(k types.T, v types.U)
	Get(k types.T)
}
