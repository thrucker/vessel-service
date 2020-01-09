package main

import (
	"errors"
	"fmt"
	"github.com/micro/go-micro"
	pb "github.com/thrucker/vessel-service/proto/vessel"
	"golang.org/x/net/context"
)

type repository interface {
	FindAvailable(specification *pb.Specification) (*pb.Vessel, error)
}

type VesselRepository struct {
	vessels []*pb.Vessel
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("no vessel found by that spec")
}

type service struct {
	repo repository
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

func main() {
	vessels := []*pb.Vessel{
		&pb.Vessel{
			Id:                   "vessel001",
			Name:                 "Boaty McBoatface",
			Capacity:             500,
			MaxWeight:            200000,
		},
	}
	repo := &VesselRepository{vessels}

	srv := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)

	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
