package main

import (
	"context"
	"errors"
	pb "github.com/thrucker/vessel-service/proto/vessel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type repository interface {
	FindAvailable(specification *pb.Specification) (*pb.Vessel, error)
	Create(vessel *pb.Vessel) error
}

type MongoRepository struct {
	collection *mongo.Collection
}

func (repo *MongoRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	filter := bson.M{
		"capacity":  bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"$gte": spec.MaxWeight},
	}
	cur, err := repo.collection.Find(context.Background(), filter)

	if err != nil {
		log.Println("error finding vessel")
		return nil, err
	}

	for cur.Next(context.Background()) {
		var vessel *pb.Vessel
		if err := cur.Decode(&vessel); err != nil {
			return nil, err
		}
		return vessel, nil
	}

	return nil, errors.New("no pb found by that spec")
}

func (repo *MongoRepository) Create(vessel *pb.Vessel) error {
	_, err := repo.collection.InsertOne(context.Background(), vessel)
	return err
}
