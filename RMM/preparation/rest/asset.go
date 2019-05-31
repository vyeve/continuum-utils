package rest

import (
	"errors"
)

// AssetClient is a REST client to interact with Asset microservice.
var AssetClient ClientInterface

const (
	assetDTPath = "http://internal-Continuum-Asset-Service-ELB-Int-1972580147.ap-south-1.elb.amazonaws.com/asset/v1"
	assetQAPath = "http://internal-Continuum-Asset-Service-ELB-QA-764891423.us-east-1.elb.amazonaws.com/asset/v1"
)

// LoadAsset creates and initializes a REST client to interact with Asset microservice.
func LoadAsset(env Environment) (err error) {
	var path string
	switch env {
	case DTEnvironment:
		path = assetDTPath
	case QAEnvironment:
		path = assetQAPath
	default:
		return errors.New("not support environment")
	}
	AssetClient, err = NewClient("Asset", path)
	if err != nil {
		return err
	}
	return nil
}
