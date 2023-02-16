package rediss

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
)

// SetToken store token in redis
func SetToken(ctx context.Context, username string, token interface{}) error {
<<<<<<< HEAD
=======
	if rdb == nil {
		initRedis()
	}
>>>>>>> 2bb4450e1ea238fa6bab7b3bbf098c30b1fd617d
	key := "token"
	if username == "" {
		return errors.New("argument is null")
	}

	err := rdb.HSet(ctx, key, username, token).Err()
	if err != nil {
		logs.Errorf("insert to redis failed: ", err.Error())
		return errors.New(err.Error())
	}
	// log.Println("HSET ", key, " ", username, " ", token)
	return nil
}
