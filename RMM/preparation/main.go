package main

import (
	"fmt"
	"log"

	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/persistency/postgres"

	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/writer"

	appLoader "github.com/vitaliyyevenko/continuum-utils/RMM/preparation/app-loader"
	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/generation"
	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/rest"
)

var env = rest.DTEnvironment

func main() {

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
