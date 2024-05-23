package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	proto_v1 "github.com/MinBZK/logboek-dataverwerkingen-logboek/libs/logboek-go/proto/logboek/v1"

	"github.com/MinBZK/logboek-dataverwerkingen-logboek/server/pkg/storage"
)

type Server struct {
	grpcServer *grpc.Server
	store      storage.Store
}

func New(store storage.Store) (*Server, error) {
	s := &Server{store: store}

	return s, nil
}

func (s *Server) Start(address string) error {
	l, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s.grpcServer = grpc.NewServer()

	if err := s.store.Initialize(); err != nil {
		return fmt.Errorf("failed to initialize storage: %v", err)
	}
	defer s.store.Close()

	proto_v1.RegisterLogboekServiceServer(s.grpcServer, LogboekService{store: s.store})

	log.Printf("Listening on %v\n", l.Addr())
	err = s.grpcServer.Serve(l)

	return err
}
