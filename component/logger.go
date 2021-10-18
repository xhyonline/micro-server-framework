package component

import (
	"github.com/xhyonline/micro-server-framework/configs"

	"github.com/xhyonline/xutil/xlog"
)

var Logger *xlog.MyLogger

func RegisterLogger() Option {
	return func() {
		if configs.Env == "dev" {
			Logger = xlog.Get().Debugger()
			return
		}
		Logger = xlog.Get().Product("./runner.log", true)
	}
}
