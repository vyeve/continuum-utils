package models

// Generator generates records
type Generator interface {
	GetProducts() (productIDs []string, err error)
	GetPartnerIDs(productIDs []string) (partnerIDs []string, err error)
	GetEndpointByID(endpointID string) (asset AssetCollection, err error)
	GetEndpoints(partnerIDs []string) (asset []AssetCollection, err error)
}
