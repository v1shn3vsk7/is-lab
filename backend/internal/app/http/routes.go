package http

func (s *Server) setupRoutes() {
	s.server.GET("/liveness", s.handlers.Liveness)

	s.server.POST("/api/signup", s.handlers.CreateUser)
	s.server.GET("/api/find_user", s.handlers.GetUser)
	s.server.GET("/api/list_users", s.handlers.ListUsers)
}
