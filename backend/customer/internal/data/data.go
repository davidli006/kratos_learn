package data

import (
	"customer/internal/conf"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
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
	redisUrl := fmt.Sprintf("redis://%s?dial_timeout=1", c.Redis.Addr)

	options, err := redis.ParseURL(redisUrl)
	if err != nil {
		data.Rdb = nil
	}
	data.Rdb = redis.NewClient(options)

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return data, cleanup, nil
}
