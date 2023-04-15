package http

func (s *Server) NewRouter() {
	s.App.POST("/user", s.handler.CreateUser)
	s.App.GET("/user/:id", s.handler.GetUserByID)

	s.App.GET("/transactions", s.handler.AllTransactions)
	s.App.POST("/transactions", s.handler.TransactionsCreate)
}
