package main

import "github.com/gorilla/mux"

//router service struct
type Server struct {
	router *mux.Router
}

//config struct
type Config struct {
	CRUDHost          string
	CRUDPort          string
	DATAEXTRACTORPort string
}

type AssetID struct {
	AssetID string `json:"id"`
}

type AssetRegisterResponse struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	SerialNo           string `json:"serialno"`
	Size               string `json:"size"`
	Type               string `json:"type"`
	Class              string `json:"class"`
	Dimension1Val      string `json:"dimension1val"`
	Dimension2Val      string `json:"dimension2val"`
	Dimension3Val      string `json:"dimension3val`
	Dimension4Val      string `json:"dimension4val"`
	Dimension5Val      string `json:"dimension5val"`
	Dimension6Val      string `json:"dimension6val"`
	Extent             string `json:"extent"`
	ExtentConfidence   string `json:"extentconfidence"`
	DeRecognitionvalue string `json:"derecognitionvalue"`
}

type AssetList struct {
	Assets []AssetRegisterResponse `json:"assets"`
}

type Application struct {
	Name               string
	Description        string
	SerialNo           string
	Size               string
	Type               string
	Class              string
	Dimension1Val      string
	Dimension2Val      string
	Dimension3Val      string
	Dimension4Val      string
	Dimension5Val      string
	Dimension6Val      string
	Extent             string
	ExtentConfidence   string
	DeRecognitionvalue string
}

type AssetListCSV struct {
	Assets []Application
}
