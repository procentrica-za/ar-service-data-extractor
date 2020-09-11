package main

import "github.com/gorilla/mux"

//router service struct
type Server struct {
	router *mux.Router
}

//config struct
type Config struct {
	CRUDHost            string
	CRUDPort            string
	DATAEXTRACTORPort string
}
