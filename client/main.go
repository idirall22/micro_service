package main

import (
	"context"
	pb "github.com/idirall22/micro_services/proto/consignment"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewShippingServiceClient(conn)

	c := &pb.Consignment{
		Id:          "01",
		Description: "this is a simple desc",
		Weight:      500,
		Containers: []*pb.Container{
			{Id: "01", CustomerId: "cust001", Origin: "origin", UserId: "1234"},
		},
		VesselId: "215",
	}

	_, err = client.CreateConsignment(context.Background(), c)

	consignments, err := client.GetConsignment(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(consignments)
}
