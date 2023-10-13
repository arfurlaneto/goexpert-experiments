package main

import (
	"context"
	"io"

	"goexpertgrpc/internal/pb"
)

type EntityService struct {
	pb.UnimplementedEntityServiceServer
	database []*pb.Entity
}

func NewEntityService() *EntityService {
	return &EntityService{
		database: make([]*pb.Entity, 0),
	}
}

func (e *EntityService) CreateEntity(ctx context.Context, req *pb.CreateEntityRequest) (*pb.CreateEntityResponse, error) {
	e.database = append(e.database, req.Entity)
	return &pb.CreateEntityResponse{
		Entity: req.Entity,
	}, nil
}

func (e *EntityService) CreateEntityStreamBidirectional(stream pb.EntityService_CreateEntityStreamBidirectionalServer) error {
	for {
		entity, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		e.database = append(e.database, entity.Entity)

		err = stream.Send(&pb.CreateEntityResponse{
			Entity: entity.Entity,
		})

		if err != nil {
			return err
		}
	}
}

func (e *EntityService) ListEntities(context.Context, *pb.ListEntitiesRequest) (*pb.ListEntitiesResponse, error) {
	return &pb.ListEntitiesResponse{
		Entities: e.database,
	}, nil
}
