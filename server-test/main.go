package main

import (
	"log"
	"net"

	"github.com/bocanada/grpc/database"
	"github.com/bocanada/grpc/server"
	"github.com/bocanada/grpc/testpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	PORT   = ":5070"
	DB_URL = "postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable"
)

func main() {
	list, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal(err)
	}
	repo, err := database.NewPostgresRepository(DB_URL)
	if err != nil {
		log.Fatal(err)
	}
	server := server.NewTestServer(repo)

	// Create a grpc server and register our StudentServer
	s := grpc.NewServer()
	testpb.RegisterTestServiceServer(s, server)

	// Use reflection to self-document
	reflection.Register(s)

	log.Println("Listening on port ", PORT)
	// Eq ListenAndServe
	if err := s.Serve(list); err != nil {
		log.Fatal(err)
	}
}
