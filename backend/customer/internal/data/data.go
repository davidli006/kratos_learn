package data

import (
	"context"
	"customer/internal/conf"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewCustomerData)

// Data .
type Data struct {
	// TODO wrapped database client
	Rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {

	data := &Data{}
	data.Rdb = createRedis(c)

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return data, cleanup, nil
}

func createRedis(c *conf.Data) *redis.Client {
	redisUrl := fmt.Sprintf("redis://%s?dial_timeout=1", c.Redis.Addr)
	options, err := redis.ParseURL(redisUrl)
	if err != nil {
		return nil
	}
	return redis.NewClient(options)
}

func createGRPC(c *conf.Data) {
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(c.Grpc.VerifyCode))
}
