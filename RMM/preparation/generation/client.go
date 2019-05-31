package generation

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"runtime"
	"strings"

	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/models"
	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/rest"
)

const chanSize = 100

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

type result struct {
	err       error
	ids       []string
	assets    []models.AssetCollection
	partnerID string
}
type client struct {
	assetClient rest.ClientInterface
	entClient   rest.ClientInterface
}

func (c client) GetProducts() (productIDs []string, err error) {
	path := rest.GetProductsPath
	resp, err := c.entClient.Get(path)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		log.Printf("response code [%d] is not correct", resp.StatusCode())
		return nil, errors.New("bad response code")
	}
	products := make([]models.Product, 0)
	err = json.Unmarshal(resp.Body(), &products)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	productIDs = getUniqueProductIDs(products)
	if len(productIDs) == 0 {
		return nil, errors.New("cannot find product ids")
	}
	return productIDs, nil
}

func setChan(ids []string, idChan chan string) {
	for _, id := range ids {
		idChan <- id
	}
	close(idChan)
}

func (c client) getParternsInChan(idChan chan string, resChan chan result) {
	for id := range idChan {
		ids, err := c.getPartnerIDs(id)
		resChan <- result{
			err: err,
			ids: ids,
		}
	}
	close(resChan)
}
func (c client) GetPartnerIDs(productIDs []string) (partnerIDs []string, err error) {
	inChan := make(chan string, chanSize)
	resChan := make(chan result, chanSize)
	go setChan(productIDs, inChan)

	go c.getParternsInChan(inChan, resChan)

	for res := range resChan {
		if res.err != nil {
			// log.Println(err)
			continue
		}
		partnerIDs = append(partnerIDs, res.ids...)
	}
	partnerIDs = getUniqueSlice(partnerIDs)
	if len(partnerIDs) == 0 {
		return nil, errors.New("cannot find parnter ids")
	}
	return partnerIDs, nil
}

func (c client) getPartnerIDs(productID string) (partnerIDs []string, err error) {
	path := strings.Replace(rest.GetPartnersByProductPath, "{productID}", productID, -1)
	resp, err := c.entClient.Get(path)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		log.Printf("response code [%d] is not correct", resp.StatusCode())
		return nil, errors.New("bad response code")
	}
	partner := new(models.PartnerEnt)
	err = json.Unmarshal(resp.Body(), partner)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	partnerIDs = make([]string, 0, len(partner.Entitlements))
	for _, ent := range partner.Entitlements {
		partnerIDs = append(partnerIDs, ent.PartnerID)
	}
	if len(partnerIDs) == 0 {
		return nil, errors.New("cannot find parnter ids")
	}
	return partnerIDs, nil
}

func (c client) GetEndpointByID(id string) (asset models.AssetCollection, err error) {
	return
}

func (c client) GetEndpoints(partnerIDs []string) (assets []models.AssetCollection, err error) {
	idChan := make(chan string, chanSize)
	resChan := make(chan result, chanSize)

	go setChan(partnerIDs, idChan)
	go c.getEndpointsInChan(idChan, resChan)

	for res := range resChan {
		if res.err != nil || len(res.assets) == 0 {
			continue
		}
		// fmt.Println("Partner:", res.partnerID)
		assets = append(assets, res.assets...)
	}
	assets = getUniqueEndpoints(assets)
	if len(assets) == 0 {
		return nil, errors.New("cannot find endpoints")
	}
	return
}

func (c client) getEndpoint(pid string) (assets []models.AssetCollection, err error) {
	path := strings.Replace(rest.GetEndpointsByPartnerIDPath, "{partnerID}", pid, -1)
	resp, err := c.assetClient.Get(path)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() == http.StatusNotFound {
		return assets, nil
	}
	if resp.StatusCode() != http.StatusOK {
		log.Printf("response code [%d] is not correct", resp.StatusCode())
		return nil, errors.New("bad response code")
	}
	err = json.Unmarshal(resp.Body(), &assets)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return assets, nil
}

func (c client) getEndpointsInChan(idChan chan string, resChan chan result) {
	for id := range idChan {
		ids, err := c.getEndpoint(id)
		resChan <- result{
			err:       err,
			assets:    ids,
			partnerID: id,
		}
	}
	close(resChan)
}

func getUniqueProductIDs(products []models.Product) []string {
	ids := make([]string, 0, len(products))
	unique := make(map[string]bool, len(products))
	for _, pr := range products {
		if !unique[pr.ProductID] {
			unique[pr.ProductID] = true
			ids = append(ids, pr.ProductID)
		}
	}
	return ids
}

func getUniqueSlice(in []string) (out []string) {
	out = make([]string, 0, len(in))
	unique := make(map[string]bool, len(in))
	for _, x := range in {
		if !unique[x] {
			unique[x] = true
			out = append(out, x)
		}
	}
	return out
}

func getUniqueEndpoints(in []models.AssetCollection) (out []models.AssetCollection) {
	out = make([]models.AssetCollection, 0, len(in))
	unique := make(map[string]bool, len(in))
	for _, x := range in {
		if !unique[x.EndpointID] {
			unique[x.EndpointID] = true
			out = append(out, x)
		}
	}
	return out
}
