package main

import (
	"fmt"
	pb "github.com/namely/arnoldb-protofiles/go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	address := os.Getenv("ARNOLDB_ADDR") + os.Getenv("ARNOLDB_PORT")
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDatastoreActionsClient(conn)

	valQuery := pb.Values{
		ObjectTypeId: otId,
		ObjectIds:    []*pb.Object{o},
		Fields:       []*pb.Field{f},
		Date:         0,
	}

	// Contact the server and print out its response.
	r, err := c.GetValues(context.Background(), valQuery)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response: %s", r)
}
