package main

import (
	"context"
	"fmt"
	"github.com/zhayt/simple-grpc/config"
	pb "github.com/zhayt/simple-grpc/pb/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"sync"
	"time"
)

var _defaultContextTime = 3 * time.Second

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

	conn, err := grpc.Dial(cfg.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("couldn't dial: %w", err)
	}
	defer conn.Close()

	c := pb.NewStudentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), _defaultContextTime)
	defer cancel()

	newStudent := createStudent()
	r, err := c.CreateStudent(ctx, &pb.CreateStudentRequest{Student: newStudent})
	if err != nil {
		return fmt.Errorf("failed to create student: %w", err)
	}

	log.Println(r)
	return nil
}

func createStudent() *pb.Student {
	return &pb.Student{Name: "Alem", Email: "alem@mail.ru", Password: "qwerty"}
}
