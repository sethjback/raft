package main

import (
	"context"
	"log"
	"net"

	"github.com/sethjback/raft/proto/raftproto"
	"google.golang.org/grpc"
)

type RaftService struct {
	raftproto.UnimplementedRaftServer
	GRPCServer *grpc.Server
}

func NewRaftService(address string) (*RaftService, error) {
	if address == "" {
		address = "127.0.0.1:5052"
	}

	r := &RaftService{
		GRPCServer: grpc.NewServer(),
	}

	raftproto.RegisterRaftServer(r.GRPCServer, r)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("raft service listening on %s\n", address)

	go func() {
		if err = r.GRPCServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	return r, nil
}

func (r *RaftService) RequestVote(ctx context.Context, in *raftproto.RequestVoteRequest) (*raftproto.RequestVoteResponse, error) {
	return &raftproto.RequestVoteResponse{}, nil
}

func (r *RaftService) AppendEntries(ctx context.Context, in *raftproto.AppendEntriesRequest) (*raftproto.AppendEntriesResponse, error) {
	return &raftproto.AppendEntriesResponse{}, nil
}

func (r *RaftService) Stop() error {
	log.Println("stopping raft service")
	r.GRPCServer.GracefulStop()
	return nil
}
