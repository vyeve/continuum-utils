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
	_ = p
	// env.Message = string(p)
	env.Message = dgKafkaMessage(asset)
	_ = c.srv.Publish(env)
}

func (c client) PublishAll(assets []models.AssetCollection) {
	for _, asset := range assets {
		if asset.PartnerID != "50000031" {
			continue
		}
		// if asset.EndpointID != "aca958aa-d6de-4e04-91c5-5e3a0d364d42" {
		// 	continue
		// }
		// if asset.EndpointID != "4d0261d4-f438-4b96-9964-d36d3c7ecb89" ||
		// 	asset.EndpointID != "fa012879-c01e-4c18-9ac4-50ac215cd03c" {
		// 	continue
		// }

		fmt.Println(asset.EndpointID)
		c.Publish(asset)
	}
}

func dgKafkaMessage(as models.AssetCollection) string {
	is := make([]models.InstalledSoftware, len(as.InstalledSoftwares))
	for i, s := range as.InstalledSoftwares {
		is[i] = models.InstalledSoftware{
			Name:        s.Name,
			InstallDate: s.InstallDate,
			Version:     s.Version,
			NameVersion: fmt.Sprintf("%s %s", s.Name, s.Version),
		}
	}
	pis, _ := json.Marshal(&is)

	// for _, n := range as.Networks {
	// 	nets = fmt.Sprintf("%q,%q", nets, n.IPv4)
	// }
	return fmt.Sprintf(`{"originDomain": "asset","id":"%s","partner":"%s","client":"%s","site":"%s","legacy_regid":"%s","asset": {"installed_software":%s, "os":"%s", "ram":%d,"service_pack":"%s", "baseboard_manufacturer": "M", "virtual_machine": "true","machine_name": "%s","machine_friendly_name": "%s", "v_pro": false,"vpro": false},"createTimeUTC": "%s",  "createdBy":"%s", "typt":"%s"}`,
		as.EndpointID,
		as.PartnerID,
		as.ClientID,
		as.SiteID,
		as.RegID,
		pis,
		fmt.Sprintf("%s %s", as.Os.Product, as.Os.ServicePack),
		15000,
		as.Os.ServicePack,
		as.Name,
		as.FriendlyName,
		as.CreateTimeUTC,
		as.CreatedBy,
		as.Type,
	)
}
