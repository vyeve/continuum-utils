package service

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func NewServer() *http.Server {
	middlewareManager := negroni.New()
	middlewareManager.Use(negroni.NewRecovery())
	negroniLogger := negroni.NewLogger()
	negroniLogger.ALogger = log.New(os.Stdout, "AssetMS\t| ", 0)

	router := mux.NewRouter().PathPrefix("/asset/v1").Subrouter()
	{
		router.HandleFunc("/partner/{partnerID}/endpoints/{endpointID}", GetAsset).Methods(http.MethodGet)
	}

	middlewareManager.Use(negroniLogger)
	middlewareManager.UseHandler(router)

	return &http.Server{
		Addr:    ":8841",
		Handler: middlewareManager,
	}

}
