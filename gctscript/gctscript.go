package gctscript

import (
	"github.com/aaabigfish/gocryptotrader/gctscript/modules"
	"github.com/aaabigfish/gocryptotrader/gctscript/wrappers/gct"
)

// Setup configures the wrapper interface to use
func Setup() {
	modules.SetModuleWrapper(gct.Setup())
}
