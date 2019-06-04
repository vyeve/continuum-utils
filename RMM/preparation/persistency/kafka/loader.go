package kafka

import (
	"encoding/json"
	"fmt"

	"github.com/ContinuumLLC/platform-common-lib/src/messaging"
	"github.com/gocql/gocql"
	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/models"
)

// Publisher describes methods to work with Kafka
type Publisher interface {
	Publish(asset models.AssetCollection)
	PublishAll(asset []models.AssetCollection)
}

// Client to work with Kafka
var Client Publisher

const (
	broker = "localhost:9092"
	topic  = "managed-endpoint-change"
)

// Load initializes kafka client
func Load() (err error) {
	conf := messaging.Config{
		Address: []string{broker},
		Topics:  []string{topic},
	}
	c := client{
		srv: messaging.NewService(conf),
	}
	Client = c
	return nil
}

type client struct {
	srv messaging.Service
}

func (c client) Publish(asset models.AssetCollection) {
	uuid := fmt.Sprintf("Transaction ID:%s,", gocql.TimeUUID().String())
	env := new(messaging.Envelope)
	env.Header = messaging.Header{
		"continuum.message.type":           {"NEWASSET"},
		"continuum.message.transaction.id": {uuid},
	}
	env.Topic = topic
	p, _ := json.Marshal(&asset.InstalledSoftwares)
	// env.Message = string(p)
	env.Message = dgKafkaMessage(asset.EndpointID, asset.PartnerID, asset.ClientID, asset.SiteID, asset.RegID, string(p))
	_ = c.srv.Publish(env)
}

func (c client) PublishAll(assets []models.AssetCollection) {
	for _, asset := range assets {
		if asset.PartnerID != "50000031" {
			continue
		}
		if asset.EndpointID != "aca958aa-d6de-4e04-91c5-5e3a0d364d42" {
			continue
		}
		fmt.Println(asset.Os)
		c.Publish(asset)
	}
}

func dgKafkaMessage(endpointID, partnerID, clientID, siteID, legacyRegID, message string) string {
	return fmt.Sprintf(`{"originDomain": "asset","id":"%s","partner":"%s","client":"%s","site":"%s","legacy_regid":"%s","asset": {"installedSoftwares":%s}}`,
		endpointID,
		partnerID,
		clientID,
		siteID,
		legacyRegID,
		message)
}
