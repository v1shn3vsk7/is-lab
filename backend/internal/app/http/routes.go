package http

func (s *Server) setupRoutes() {
	s.server.GET("/liveness", s.handlers.Liveness)

	s.server.POST("/api/signup", s.handlers.CreateUser)
}
