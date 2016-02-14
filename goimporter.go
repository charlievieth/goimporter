// +build !go1.5

// Package importer provides access to export data importers.
package goimporter

import (
	"runtime"

	"git.vieth.io/goimporter/internal/gccgoimporter"
	"git.vieth.io/goimporter/internal/gcimporter"
	"git.vieth.io/goimporter/internal/types"
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
