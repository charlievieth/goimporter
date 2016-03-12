// +build !go1.5

package goimporter

import (
	"testing"

	"git.vieth.io/goimporter/internal/types"
)

func BenchmarkImporter(b *testing.B) {
	m := Default().(*gcimports)
	n := 0
	for i := 0; i < b.N; i++ {
		if _, err := m.Import(packageNames[n%len(packageNames)]); err != nil {
			b.Fatal(err)
		}
		n++
		b.StopTimer()
		m.pkgs = make(map[string]*types.Package)
		b.StartTimer()
	}
}
