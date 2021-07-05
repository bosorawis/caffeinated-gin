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

func (s *Server) handleHelloWorld()func(c *gin.Context) {
	return func(c *gin.Context){
		ip, _ := c.RemoteIP()
		c.String(http.StatusOK, "Hello %s ", ip.String())
	}
}

func (s *Server) handlePostReview() func(c *gin.Context) {
	type request struct {
		Name         string   `json:"name"`
		Score        float64  `json:"score"`
		TestingNotes []string `json:"tasting-notes"`
	}
	return func(c *gin.Context){
		var data request
		c.BindJSON(&data)
		c.String(http.StatusOK, "Hello %v", data)
	}
}
