package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "" // Import the generated protobuf code

	"google.golang.org/grpc"
)

type noteServiceServer struct{}

func (s *noteServiceServer) List(ctx context.Context, empty *pb.Empty) (*pb.NoteList, error) {
	// This is where you would implement the logic to fetch and return a list of notes
	// For the sake of this example, we'll return a hard-coded list of notes
	notes := []*pb.Note{
		{Id: "1", Title: "Note 1", Content: "This is the content of note 1"},
		{Id: "2", Title: "Note 2", Content: "This is the content of note 2"},
		{Id: "3", Title: "Note 3", Content: "This is the content of note 3"},
	}
	return &pb.NoteList{Notes: notes}, nil
}

func main() {
	// Create a new gRPC server
	server := grpc.NewServer()
	fmt.Println("Starting server...")

	// Register our note service with the server
	pb.RegisterNoteServiceServer(server, &noteServiceServer{})

	// Listen for incoming connections on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Server listening on port 50051...")

	// Serve requests using the gRPC server
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
