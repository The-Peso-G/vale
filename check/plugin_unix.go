// +build !windows

package check

func loadPlugins(mgr Manager) Manager {
	return mgr
}

/* TODO: Is this even worth it?
import (
	"path/filepath"
	"plugin"
	"strings"

	"github.com/errata-ai/vale/core"
)

func loadPlugins(mgr Manager) Manager {
	loc := filepath.Join(mgr.Config.StylesPath, "plugins", "*.so")

	plugins, err := filepath.Glob(loc)
	if err != nil {
		panic(err)
	}

	for _, filename := range plugins {
		f, _ := plugin.Open(filename)

		base := filepath.Base(filename)
		symb := strings.TrimSuffix(base, filepath.Ext(base))
		name := "plugins." + symb

		symbol, err := f.Lookup(symb)
		if err != nil {
			panic(err)
		}
		p := symbol.(func() core.Plugin)()

		chk := Check{Rule: p.Rule, Scope: core.Selector{p.Scope}}
		if l, ok := mgr.Config.RuleToLevel[name]; ok {
			chk.Level = core.LevelToInt[l]
		} else {
			chk.Level = core.LevelToInt[p.Level]
		}
		mgr.AllChecks[name] = chk
	}

	return mgr
}*/
