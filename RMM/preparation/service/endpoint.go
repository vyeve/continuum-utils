package service

import (
	"net/http"

	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/persistency/postgres"
)

// GetAsset returns asset from postgres DB
func GetAsset(w http.ResponseWriter, r *http.Request) {
	partnerID := GetPartnerID(r)
	endpointID := GetEndpointID(r)
	if partnerID == "" || endpointID == "" {
		http.Error(w, "not correct data", http.StatusBadRequest)
		return
	}
	asset, err := postgres.Client.GetAsset(partnerID, endpointID)
	if err != nil {
		if err == postgres.ErrNotFound {
			SendError(w, http.StatusNotFound, err.Error())
			return
		}
		SendInternalServerError(w, err.Error())
		return
	}
	RenderJSON(w, asset)
}
