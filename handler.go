package main

import (
	pb "github.com/thrucker/vessel-service/proto/vessel"
	"golang.org/x/net/context"
	"log"
)

type handler struct {
	repository
}

func (h *handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	log.Println("Trying to find vessel")
	vessel, err := h.repository.FindAvailable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

func (h *handler) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	if err := h.repository.Create(req); err != nil {
		return err
	}
	res.Vessel = req
	res.Created = true
	return nil
}
