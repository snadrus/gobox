package pprof

import (
	"net/http"
	"net/http/pprof"

	"github.com/gin-gonic/gin"
)

// Register the standard HandlerFuncs from the net/http/pprof package with
// the provided gin.RouterGroup
func Register(r *gin.RouterGroup) {
	r = r.Group("/debug/pprof")
	r.GET("/", pprofHandler(pprof.Index))
	r.GET("/cmdline", pprofHandler(pprof.Cmdline))
	r.GET("/profile", pprofHandler(pprof.Profile))
	r.POST("/symbol", pprofHandler(pprof.Symbol))
	r.GET("/symbol", pprofHandler(pprof.Symbol))
	r.GET("/trace", pprofHandler(pprof.Trace))
	r.GET("/allocs", pprofHandler(pprof.Handler("allocs").ServeHTTP))
	r.GET("/block", pprofHandler(pprof.Handler("block").ServeHTTP))
	r.GET("/goroutine", pprofHandler(pprof.Handler("goroutine").ServeHTTP))
	r.GET("/heap", pprofHandler(pprof.Handler("heap").ServeHTTP))
	r.GET("/mutex", pprofHandler(pprof.Handler("mutex").ServeHTTP))
	r.GET("/threadcreate", pprofHandler(pprof.Handler("threadcreate").ServeHTTP))
}

func pprofHandler(h http.HandlerFunc) gin.HandlerFunc {
	handler := http.HandlerFunc(h)
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
