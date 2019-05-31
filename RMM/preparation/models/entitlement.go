package models

// Product ...
type Product struct {
	ProductID string `json:"id"`
}

// PartnerEnt ...
type PartnerEnt struct {
	Entitlements []Entitlement `json:"entitlements"`
}

// Entitlement ...
type Entitlement struct {
	PartnerID string `json:"partnerID"`
	ClientID  string `json:"clientID"`
	SiteID    string `json:"siteID"`
}
