package database

import "github.com/weaviate/weaviate-go-client/v4/weaviate"

type StorageInterface interface {
}

type Storage struct {
	Client *weaviate.Client
}
