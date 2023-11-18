package shared

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

type RedisDTO struct {
	Sender, Receiver, Message string
}

func NewRedisRepo(uri string) RedisRepo {
	client := redis.NewClient(&redis.Options{
		Addr: uri,
		DB:   0,
	})
	return RedisRepo{
		client: client,
	}
}

type RedisRepo struct {
	client *redis.Client
}

func (rr *RedisRepo) Store(dto RedisDTO) error {
	marshalledMessage, err := json.Marshal(dto)
	if err != nil {
		return err
	}

	fmt.Println(marshalledMessage)
	//Use redis client to store object
	return fmt.Errorf("unimplemented")
}

func (rr *RedisRepo) Get(sender, receiver string) ([]RedisDTO, error) {
	//use redis client to retrieve object
	return nil, fmt.Errorf("unimplemented")
}
