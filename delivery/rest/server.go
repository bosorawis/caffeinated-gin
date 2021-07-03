package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	r := gin.Default()
	srv := &Server{
		router: r,
	}
	srv.routes()
	return srv
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	 s.router.ServeHTTP(w, r)
}

func (s *Server) handleHelloWorld(c *gin.Context){
	ip, _ := c.RemoteIP()
	c.String(http.StatusOK, "Hello %s ", ip.String())
}