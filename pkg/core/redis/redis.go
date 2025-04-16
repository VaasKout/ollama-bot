package redis

import (
	"context"
	"fmt"
	redisDb "github.com/redis/go-redis/v9"
	"ollama-bot/configs"
)

type Client interface {
	SetData(key string, value string) error
	GetData(key string) string
	DeleteData(key string) error
	SAdd(key string, member string) error
	SMembers(key string) []string
	SISMembers(key string, value string) bool
	SRem(key string, member string) error
	RPush(key string, value string) error
	LPush(key string, value string) error
	LPop(key string) string
	GetSize(key string) int64
	LTrim(key string, startIndex int64) error
	LRange(key string) []string
}

type ClientImpl struct {
	redisClient *redisDb.Client
	context     context.Context
}

func New(cfg *configs.Config) Client {
	ctx := context.Background()
	client := getClient(
		cfg.RedisProps.RedisAddress,
		cfg.RedisProps.RedisUser,
		cfg.RedisProps.RedisPassword,
	)
	err := client.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}

	return &ClientImpl{
		redisClient: client,
		context:     ctx,
	}
}

func getClient(
	address string,
	username string,
	password string,
) *redisDb.Client {
	return redisDb.NewClient(&redisDb.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       0,
	})
}

func (adapter *ClientImpl) SetData(key string, value string) error {
	err := adapter.redisClient.Set(adapter.context, key, value, -1).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (adapter *ClientImpl) GetData(key string) string {
	result, err := adapter.redisClient.Get(adapter.context, key).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return result
}

func (adapter *ClientImpl) DeleteData(key string) error {
	err := adapter.redisClient.Del(adapter.context, key).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (adapter *ClientImpl) SAdd(key string, member string) error {
	err := adapter.redisClient.SAdd(adapter.context, key, member).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (adapter *ClientImpl) SMembers(key string) []string {
	result, err := adapter.redisClient.SMembers(adapter.context, key).Result()
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	return result
}

func (adapter *ClientImpl) SISMembers(key string, value string) bool {
	result, err := adapter.redisClient.SIsMember(adapter.context, key, value).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return result
}

func (adapter *ClientImpl) SRem(key string, member string) error {
	_, err := adapter.redisClient.SRem(adapter.context, key, member).Result()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

func (adapter *ClientImpl) LPush(key string, value string) error {
	err := adapter.redisClient.LPush(adapter.context, key, value).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (adapter *ClientImpl) RPush(key string, value string) error {
	err := adapter.redisClient.RPush(adapter.context, key, value).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (adapter *ClientImpl) LPop(key string) string {
	result, err := adapter.redisClient.LPop(adapter.context, key).Result()
	if err != nil {
		return ""
	}
	return result
}

func (adapter *ClientImpl) GetSize(key string) int64 {
	result, err := adapter.redisClient.LLen(adapter.context, key).Result()
	if err != nil {
		return 0
	}
	return result
}

func (adapter *ClientImpl) LRange(key string) []string {
	result, err := adapter.redisClient.LRange(adapter.context, key, 0, 100).Result()
	if err != nil {
		fmt.Println(err)
		return []string{}
	}

	return result
}

func (adapter *ClientImpl) LTrim(key string, startIndex int64) error {
	_, err := adapter.redisClient.LTrim(adapter.context, key, startIndex, -1).Result()
	if err != nil {
		fmt.Println(err)
	}
	return err
}
