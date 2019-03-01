package service

// Add routes to the server's router (ServeMux)
func (s *server) AddRoutes() {
  s.router.HandleFunc("/redirect", s.redirect())
}
