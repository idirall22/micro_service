package main

import (
	"context"
	"log"
	"sync"

	pb "github.com/idirall22/micro_services/proto/consignment"
	"github.com/micro/go-micro"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository struct
type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

// Create create consignment
func (r *Repository) Create(con *pb.Consignment) (*pb.Consignment, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.consignments = append(r.consignments, con)

	return con, nil
}

// GetAll return consignments
func (r *Repository) GetAll() []*pb.Consignment {
	return r.consignments
}

type service struct {
	repo repository
}

// CreateConsignment create
func (s *service) CreateConsignment(ctx context.Context, c *pb.Consignment, res *pb.Reponse) error {
	con, err := s.repo.Create(c)
	res.Created = true
	res.Consignment = con
	return err
}

// GetConsignment get
func (s *service) GetConsignment(ctx context.Context, req *pb.GetRequest, res *pb.Reponse) error {
	res.Consignments = s.repo.GetAll()
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)

	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo: &Repository{}})
	log.Fatal(srv.Run())
}
