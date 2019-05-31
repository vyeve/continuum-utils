package generation

import (
	"errors"

	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/models"
	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/rest"
)

// Client ...
var Client models.Generator

// Load ...
func Load() (err error) {
	if rest.AssetClient == nil || rest.EntitlementClient == nil {
		return errors.New("rest client is not initialized")
	}
	cl := client{
		assetClient: rest.AssetClient,
		entClient:   rest.EntitlementClient,
	}

	Client = cl
	return nil
}
