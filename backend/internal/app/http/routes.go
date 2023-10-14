package http

func (s *Server) setupRoutes() {
	s.server.GET("/liveness", s.handlers.Liveness)

	//s.server.POST("/api/signup", s.handlers.CreateUser)
	s.server.GET("/api/find_user", s.handlers.GetUser)
	s.server.GET("/api/list_users", s.handlers.ListUsers)
	s.server.POST("/api/update_user_password", s.handlers.UpdateUserPassword)
	s.server.PATCH("/api/update_user", s.handlers.UpdateUser)
	s.server.POST("/api/setup_password", s.handlers.SetupPassword)
	s.server.POST("/api/create_user", s.handlers.CreateUserFromAdmin)
}
