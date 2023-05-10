package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/zhayt/simple-grpc/config"
	pb "github.com/zhayt/simple-grpc/pb/user_v1"
	"github.com/zhayt/simple-grpc/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"sync"
	"time"
)

const grpcPort = "5001"

type Server struct {
	pb.UnimplementedStudentServiceServer
	storage storage.IStorage
}

func NewServer(storage storage.IStorage) *Server {
	return &Server{storage: storage}
}

func (s *Server) CreateStudent(ctx context.Context, req *pb.CreateStudentRequest) (*pb.CreateStudentResponse, error) {
	log.Println("Create user")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req.Student.Id = randomID()

	res, err := s.storage.CreateStudent(ctx, req.Student)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "couldn't create user: %w", err)
	}

	log.Printf("Student created with id %v", res)

	return &pb.CreateStudentResponse{Id: req.Student.Id}, nil
}

func randomID() string {
	return uuid.New().String()
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	var once sync.Once

	once.Do(config.MustPrepareEnv)

	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	client, err := storage.Dial(cfg)
	if err != nil {
		return err
	}

	repo := storage.NewStorage(client)

	srv := NewServer(repo)

	lis, err := net.Listen("tcp", net.JoinHostPort("", cfg.Port))
	if err != nil {
		return fmt.Errorf("couldn't start listen: %w", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterStudentServiceServer(grpcServer, srv)

	log.Printf("server listening at %s", lis.Addr())
	if err = grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("couldb't serve grpc server: %w", err)
	}

	return nil
}
