package main

import (
	"log"
	"net/http"
	handler "sharding-test-golang/internal/api"
	"sharding-test-golang/internal/storage/router"
	"sharding-test-golang/internal/storage/shards"
	"sharding-test-golang/internal/user"

	_ "github.com/lib/pq"
)

func main() {
	// Create shards
	shard0, err := shards.NewShard("shard0", "postgres://user:pass@localhost:5433/shard0?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	shard1, err := shards.NewShard("shard1", "postgres://user:pass@localhost:5434/shard1?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	shard2, err := shards.NewShard("shard2", "postgres://user:pass@localhost:5435/shard2?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Linking shards to router
	shardsList := []router.UserStorage{shard0, shard1, shard2}
	router := router.NewShardRouter(shardsList)

	service := user.NewUserService(router)

	h := handler.NewHandler(service)
	mux := http.NewServeMux()
	h.RegisterRoutes(mux)

	log.Println("Servidor escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", mux)

}
