package app

func (s *server) router() {
	s.app.Post("/send-code", s.handlers.SendCode)
	s.app.Post("/check-code", s.handlers.CheckCode)
	s.app.Get("/healthz", s.handlers.HealthCheck)
}
