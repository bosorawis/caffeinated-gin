package rest


func (s *Server) routes(){
	s.router.GET("/",  s.handleHelloWorld)
}
