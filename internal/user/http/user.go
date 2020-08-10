package http

import "github.com/gin-gonic/gin"

func (s *Server) users(g *gin.Context) {
	g.JSON(200, map[string]interface{}{
		"name": "wsc",
	})

}
