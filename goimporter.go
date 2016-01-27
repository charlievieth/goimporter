// +build !go1.5

package goimporter

import (
	"runtime"

	"git.vieth.io/goimporter/vendor/gccgoimporter"
	"git.vieth.io/goimporter/vendor/gcimporter"
	"git.vieth.io/goimporter/vendor/types"
)

func For(compiler string) types.Importer {
	switch compiler {
	case "gc":
		return gcimporter.Import
	case "gccgo":
		var inst gccgoimporter.GccgoInstallation
		if err := inst.InitFromDriver("gccgo"); err != nil {
			return nil
		}
		return inst.GetImporter(nil, nil)
	}
	// compiler not supported
	return nil
}

func Default() types.Importer {
	return For(runtime.Compiler)
}
