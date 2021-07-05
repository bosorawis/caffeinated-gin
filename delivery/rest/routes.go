package rest


func (s *Server) routes(){
	s.router.GET("/",  s.handleHelloWorld)
	s.router.POST("/post",  s.handlePostReview())
}
