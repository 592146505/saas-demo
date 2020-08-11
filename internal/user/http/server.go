package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"runtime"
	"saas-demo/common/conf"
	"strconv"
	"time"
)

// Server is http server.
type Server struct {
	Engine *gin.Engine
	Config *conf.HTTPServerConfig
}

// New new a http server.
func New(c *conf.HTTPServerConfig) *Server {
	engine := gin.New()
	engine.Use(recoverHandler)

	s := &Server{
		Engine: engine,
		Config: c,
	}
	s.initRouter()
	return s
}

func (s *Server) initRouter() {
	group := s.Engine.Group("/api/v1")
	group.GET("/users", s.users) //频道列表
}

func (s *Server) Run() {
	go func() {
		if err := s.Engine.Run(":" + strconv.FormatUint(s.Config.Port, 10)); err != nil {
			panic(err)
		}
	}()
}

// Close close the server.
func (s *Server) Close() {

}

func recoverHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			httprequest, _ := httputil.DumpRequest(c.Request, false)
			pnc := fmt.Sprintf("[Recovery] %s panic recovered:\n%s\n%s\n%s", time.Now().Format("2006-01-02 15:04:05"), string(httprequest), err, buf)
			fmt.Print(pnc)
			c.AbortWithStatus(500)
		}
	}()
	c.Next()
}
