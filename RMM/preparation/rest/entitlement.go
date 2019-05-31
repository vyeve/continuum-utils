package rest

import (
	"errors"
)

// EntitlementClient is a REST client to interact with Entitlement microservice.
var EntitlementClient ClientInterface

const (
	// entitlementDTPath = "http://internal-entitlement-partnerspecific-72155629.ap-south-1.elb.amazonaws.com/entitlement/v1"
	entitlementDTPath = "http://internal-entitlement-partneragnostic-1185276102.ap-south-1.elb.amazonaws.com/entitlement/v1"
	entitlementQAPath = "http://internal-Entitlement-PartnerAgnostic-QA-297548695.us-east-1.elb.amazonaws.com/entitlement/v1"
	// entitlementQAPath = "http://internal-Entitlement-PartnerSpecific-QA-2003801024.us-east-1.elb.amazonaws.com/entitlement/v1"
)

// LoadEntitlement creates and initializes a REST client to interact with Entitlement microservice.
func LoadEntitlement(env Environment) (err error) {
	var path string
	switch env {
	case DTEnvironment:
		path = entitlementDTPath
	case QAEnvironment:
		path = entitlementQAPath
	default:
		return errors.New("not support environment")
	}
	EntitlementClient, err = NewClient("Entitlement", path)
	if err != nil {
		return err
	}
	return nil
}
