package service

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/persistency/postgres"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetAsset returns asset from postgres DB
func GetAsset(w http.ResponseWriter, r *http.Request) {
	partnerID := GetPartnerID(r)
	endpointID := GetEndpointID(r)
	if partnerID == "" || endpointID == "" {
		http.Error(w, "not correct data", http.StatusBadRequest)
		return
	}
	// qaAsset, err := postgres.Client.GetByPartner("50000031")
	// if err != nil {
	// 	SendError(w, http.StatusNotFound, err.Error())
	// 	return
	// }
	// n := rand.Intn(len(qaAsset))
	// fmt.Println("Len:", len(qaAsset), "rand:", n)
	asset, err := postgres.Client.GetAsset(partnerID, endpointID)
	if err != nil {
		if err == postgres.ErrNotFound {
			SendError(w, http.StatusNotFound, err.Error())
			return
		}
		SendInternalServerError(w, err.Error())
		return
	}
	// if endpointID == "20fe5869-f0dc-4970-8b24-6598f35cfb4b" {
	// 	SendInternalServerError(w, "not found")
	// 	return
	// }
	// qa := qaAsset[n]
	// asset.Networks = qa.Networks
	// asset.Processors = qa.Processors
	// asset.RaidController = qa.RaidController
	// asset.InstalledSoftwares = qa.InstalledSoftwares
	// asset.Keyboards = qa.Keyboards
	// asset.Monitors = qa.Monitors
	// asset.Mouse = qa.Mouse
	// asset.Memory = qa.Memory
	// asset.PhysicalDrives = qa.PhysicalDrives
	// asset.Shares = qa.Shares
	// asset.Users = qa.Users
	// asset.SoftwareLicenses = qa.SoftwareLicenses
	// asset.Bios = qa.Bios
	// asset.Drives = qa.Drives
	// asset.PhysicalDrives = qa.PhysicalDrives

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
			fmt.Println("Before:", as.PartnerID)
			ids = append(ids, strings.Replace(as.PartnerID, "\"", "", -1))
		}
	}
	fmt.Println("All partners:", ids)
	RenderJSON(w, ids)
}

func GetSites(w http.ResponseWriter, r *http.Request) {

	partnerID := GetPartnerID(r)
	if partnerID == "" {
		http.Error(w, "not correct data", http.StatusBadRequest)
		return
	}
	fmt.Println()
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
