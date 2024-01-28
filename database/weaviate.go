package database

import (
	"os"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
)

func NewClient() (s *Storage) {
	cfg := weaviate.Config{
		Host:   os.Getenv("HOST"),
		Scheme: "http",
	}
	client, err := weaviate.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	return &Storage{
		Client: client,
	}
}

// 	schema, err := client.Schema().Getter().Do(context.Background())
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%v", schema)
// }
