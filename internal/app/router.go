package app

func (s *server) router() {
	s.app.Post("/send-code", s.handlers.SendCode)
	s.app.Post("/check-code", s.handlers.CheckCode)
	s.app.Get("/healthz", s.handlers.HealthCheck)
	s.app.Post("/register", s.handlers.SaveUser)
	s.app.Put("/update-profile", s.handlers.UpdateUser)
	s.app.Get("/get-profile/:student_id", s.handlers.GetUser)
}
