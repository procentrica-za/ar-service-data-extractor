package main

//create routes
func (s *Server) routes() {
	s.router.HandleFunc("/extract", s.handleextractassets()).Methods("GET")
}
