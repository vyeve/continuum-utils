package service

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// New returns new HTTP server
func New(port int) *http.Server {
	middlewareManager := negroni.New()
	middlewareManager.Use(negroni.NewRecovery())
	negroniLogger := negroni.NewLogger()
	negroniLogger.ALogger = log.New(os.Stdout, "AssetMS\t| ", 0)

	router := mux.NewRouter().PathPrefix("/asset/v1").Subrouter()
	{
		router.HandleFunc("/partner/{partnerID}/endpoints/{endpointID}", GetAsset).Methods(http.MethodGet)
		router.HandleFunc("/partner/{partnerID}", GetAssetsByPartner).Methods(http.MethodGet)
		router.HandleFunc("/partner/0/partners", GetPartners).Methods(http.MethodGet)
		router.HandleFunc("/partners/{partnerID}/sites", GetSites).Methods(http.MethodGet)
		router.HandleFunc("/partner/{partnerID}/sites/{siteID}/summary", GetAssetsByPartnerAndSite).Methods(http.MethodGet)
	}

	middlewareManager.Use(negroniLogger)
	middlewareManager.UseHandler(router)

	fmt.Printf("Microservice starts listening on port: %d\n", port)
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: middlewareManager,
	}

}
