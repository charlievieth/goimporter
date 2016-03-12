// +build !go1.5

package lru

import (
	"os"

	"git.vieth.io/goimporter/internal/types"
)

type Package struct {
	Pkg  *types.Package
	Info os.FileInfo
}
