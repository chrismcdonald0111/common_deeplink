package service

func (s *server) AddRoutes() {
  s.router.HandleFunc("/red", s.redirect())
}
