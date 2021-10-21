package internal

import (
	"net/http"
	_ "net/http/pprof"
	"os"

	. "github.com/xhyonline/micro-server-framework/component"
)

// initDebugPProf 内部启动 pprof
func initDebugPProf() {
	go func() {
		if err := http.ListenAndServe(internalAddress(), nil); err != nil {
			Logger.Errorf("pprof 服务启动失败")
			os.Exit(1)
		}
	}()
}
