package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"database/sql"

	_ "github.com/lib/pq"

	appLoader "github.com/vitaliyyevenko/continuum-utils/RMM/preparation/app-loader"
	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/generation"
	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/models"
	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/persistency/postgres"
	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/rest"
	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/service"
	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/writer"
)

var env = rest.DTEnvironment

func prepareDB(conf models.Configuration) (err error) {
	if !conf.SetUpDB {
		return
	}

	var (
		ids, partners []string
		endpoints     []models.AssetCollection
	)
	ids, err = generation.Client.GetProducts()
	if err != nil {
		return err
	}
	fmt.Println("Length Products:", len(ids))
	partners, err = generation.Client.GetPartnerIDs(ids)
	if err != nil {
		return err
	}
	fmt.Println("Length Partners:", len(partners))
	endpoints, err = generation.Client.GetEndpoints(partners)
	if err != nil {
		return err
	}
	fmt.Println("Length Endpoints:", len(endpoints))
	return postgres.Client.WriteAll(endpoints)
}

func mainZ() {
	config := models.Configuration{
		SetUpDB:     true,
		Port:        8841,
		Environment: rest.QAEnvironment,
	}
	// Load clients
	err := appLoader.Load(config.Environment)
	if err != nil {
		log.Fatal(err)
	}
	err = prepareDB(config)
	if err != nil {
		log.Fatal(err)
	}
	srv := service.New(config.Port)
	if err = srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func main2() {

	fmt.Println(env.String())
	err := appLoader.Load(env)
	if err != nil {
		log.Fatal(err)
	}
	ids, err := generation.Client.GetProducts()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Length Products:", len(ids))
	partners, err := generation.Client.GetPartnerIDs(ids[:20])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Length Partners:", len(partners))

	endpoints, err := generation.Client.GetEndpoints(partners[:20])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Length Endpoints:", len(endpoints))
	err = writer.Client.WriteAll(endpoints)
	if err != nil {
		log.Println(err)
	}
	for _, en := range endpoints {
		err = postgres.Client.Write(en)
		if err != nil {
			log.Println(err)
		}
	}
}

const (
	dtFile = "./endpointsDT.txt"
	qaFile = "./endpointsQA.txt"
)

func main() {
	fmt.Println("Lets go...")
	err := appLoader.Load(env)
	if err != nil {
		log.Fatal(err)
	}
	assets, err := readFile(dtFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Assets:", len(assets))
	for _, as := range assets {
		// if as.BaseBoard.HardwareUUID != "" {
		// 	fmt.Printf("Endpoint: %s; BB:%+v\n", as.EndpointID, as.BaseBoard)
		// }
		if len(as.Firewall) != 0 {
			fmt.Printf("Endpoint: %s\n", as.EndpointID)
		}
		// for _, m := range as.Networks {
		// 	if m.IPv4IntfCname != "" {
		// 		fmt.Printf("Endpoint: %s\n", as.EndpointID)
		// 		break
		// 	}
		// }
	}
	// err = postgres.Client.WriteAll(assets)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func readFile(fileName string) (assets []models.AssetCollection, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	var line string
	for {
		line, err = reader.ReadString('\n')

		if err == nil || err == io.EOF {
			if len(line) > 1 {
				ass := models.AssetCollection{}
				errM := json.Unmarshal([]byte(line), &ass)
				if errM == nil {
					assets = append(assets, ass)
				}
			}
		}

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		return nil, err
	}

	return assets, nil
}

func main5() {

	connStr := "user=postgres dbname=dg sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from asset where endpointID='2951218b-2bbe-4f3a-90d1-ead5c396a426'")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	assets := []AssetRaw{}

	for rows.Next() {
		p := AssetRaw{}
		err := rows.Scan(&p.EndpointID, &p.PartnerID, &p.RawAsset)
		if err != nil {
			fmt.Println(err)
			continue
		}
		assets = append(assets, p)
	}
	for _, p := range assets {
		fmt.Println(p.EndpointID, p.PartnerID)
		fmt.Println(p.RawAsset)
	}
}

type AssetRaw struct {
	EndpointID string `json:"endpointID"`
	PartnerID  string `json:"partnerID"`
	RawAsset   string `json:"rawAsset"`
}
