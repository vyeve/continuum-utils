package service

import (
	"net/http"
	"strconv"

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

// GetAssetsByPartner returns assets from postgres DB
func GetAssetsByPartner(w http.ResponseWriter, r *http.Request) {
	partnerID := GetPartnerID(r)
	if partnerID == "" {
		http.Error(w, "not correct data", http.StatusBadRequest)
		return
	}
	assets, err := postgres.Client.GetByPartner(partnerID)
	if err != nil {
		if err == postgres.ErrNotFound {
			SendError(w, http.StatusNotFound, err.Error())
			return
		}
		SendInternalServerError(w, err.Error())
		return
	}
	RenderJSON(w, assets)
}

// GetAllAssets returns all assets from postgres DB
// Don't use !!!
func GetAllAssets(w http.ResponseWriter, r *http.Request) {
	assets, err := postgres.Client.GetAll()
	if err != nil {
		if err == postgres.ErrNotFound {
			SendError(w, http.StatusNotFound, err.Error())
			return
		}
		SendInternalServerError(w, err.Error())
		return
	}
	RenderJSON(w, assets)
}

func GetPartners(w http.ResponseWriter, r *http.Request) {
	assets, err := postgres.Client.GetAll()
	if err != nil {
		SendError(w, http.StatusNotFound, err.Error())
		return
	}

	ids := make([]string, 0, len(assets))
	unique := make(map[string]bool)
	for _, as := range assets {
		if !unique[as.PartnerID] {
			unique[as.PartnerID] = true
			ids = append(ids, as.PartnerID)
		}
	}
	RenderJSON(w, ids)
}

func GetSites(w http.ResponseWriter, r *http.Request) {
	partnerID := GetPartnerID(r)
	if partnerID == "" {
		http.Error(w, "not correct data", http.StatusBadRequest)
		return
	}
	assets, err := postgres.Client.GetByPartner(partnerID)
	if err != nil {
		if err == postgres.ErrNotFound {
			SendError(w, http.StatusNotFound, err.Error())
			return
		}
		SendInternalServerError(w, err.Error())
		return
	}
	unique := make(map[string]bool)
	data := outdataSites{}
	for _, as := range assets {
		if !unique[as.SiteID] {
			unique[as.SiteID] = true
			id, err := strconv.Atoi(as.SiteID)
			if err != nil {
				continue
			}
			data.Sites = append(data.Sites, Site{
				ID: id,
			})
		}

	}
	RenderJSON(w, data)
}

type Site struct {
	ID int `json:"siteId"`
}

// outdata is the internal type for retrieving only sites
type outdataSites struct {
	Sites []Site `json:"outdata"`
}


// GetAssetsByPartner returns assets from postgres DB
func GetAssetsByPartnerAndSite(w http.ResponseWriter, r *http.Request) {
	partnerID := GetPartnerID(r)
	if partnerID == "" {
		http.Error(w, "not correct data", http.StatusBadRequest)
		return
	}
	assets, err := postgres.Client.GetByPartner(partnerID)
	if err != nil {
		if err == postgres.ErrNotFound {
			SendError(w, http.StatusNotFound, err.Error())
			return
		}
		SendInternalServerError(w, err.Error())
		return
	}
	RenderJSON(w, assets)
}
