package storage

import (
	"context"
	"fmt"
	"github.com/zhayt/simple-grpc/config"
	pb "github.com/zhayt/simple-grpc/pb/user_v1"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type IStorage interface {
	CreateStudent(ctx context.Context, student *pb.Student) (interface{}, error)
}

type Storage struct {
	collection *mongo.Collection
}

func NewStorage(client *mongo.Client) *Storage {
	collection := client.Database("student").Collection("students")

	return &Storage{collection: collection}
}

func Dial(cfg *config.Config) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.DBConnection))
	if err != nil {
		return nil, fmt.Errorf("couldn't get client: %w", err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("couldn't connect to db: %w", err)
	}

	return client, nil
}

func (r *Storage) CreateStudent(ctx context.Context, student *pb.Student) (interface{}, error) {
	res, err := r.collection.InsertOne(ctx, student)
	if err != nil {
		return "", fmt.Errorf("couldn't insert student: %w", err)
	}

	return res.InsertedID, nil
}
