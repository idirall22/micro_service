package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "github.com/idirall22/micro_services/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

func (s *service) CreateConsignment(ctx context.Context, con *pb.Consignment) (*pb.Reponse, error) {
	c, err := s.repo.Create(con)
	if err != nil {
		return nil, err
	}
	return &pb.Reponse{Created: true, Consignment: c}, nil
}

func (s *service) GetConsignment(ctx context.Context, req *pb.GetRequest) (*pb.Reponse, error) {
	return &pb.Reponse{Consignments: s.repo.GetAll()}, nil
}

func main() {
	l, err := net.Listen("tcp", ":5000")

	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()

	pb.RegisterShippingServiceServer(srv, &service{repo: &Repository{}})
	reflection.Register(srv)

	log.Println("Runing on port 5000")
	log.Fatal(srv.Serve(l))
}
