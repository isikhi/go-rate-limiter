package redis

import (
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"

	"github.com/isikhi/go-rate-limiter/config"
)

func New(cfg config.Cache) *redis.Client {
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	log.Printf("redis is connecting to %v\n", addr)
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Pass,
		DB:       cfg.Name,
	})
}

func NewCluster(cfg config.Cache) *redis.ClusterClient {
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: cfg.Hosts,

		// To route commands by latency or randomly, enable one of the following.
		RouteByLatency: true,
		//RouteRandomly: true,
	})
}
