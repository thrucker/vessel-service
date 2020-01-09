package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	pb "github.com/thrucker/vessel-service/proto/vessel"
	"log"
	"os"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	srv := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)

	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	vesselCollection := client.Database("shippy").Collection("vessels")

	repository := &MongoRepository{vesselCollection}
	h := &handler{repository}

	pb.RegisterVesselServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
