package main

import (
	"context"
	"errors"
	"github.com/thrucker/vessel-service/proto/vessel"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository interface {
	FindAvailable(specification *vessel.Specification) (*vessel.Vessel, error)
	Create(vessel *vessel.Vessel) error
}

type MongoRepository struct {
	collection *mongo.Collection
}

func (repo *MongoRepository) FindAvailable(spec *vessel.Specification) (*vessel.Vessel, error) {
	cur, err := repo.collection.Find(context.Background(), nil, nil)

	if err != nil {
		return nil, err
	}

	for cur.Next(context.Background()) {
		var vessel *vessel.Vessel
		if err := cur.Decode(&vessel); err != nil {
			return nil, err
		}
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}

	return nil, errors.New("no vessel found by that spec")
}

func (repo *MongoRepository) Create(vessel *vessel.Vessel) error {
	_, err := repo.collection.InsertOne(context.Background(), vessel)
	return err
}
