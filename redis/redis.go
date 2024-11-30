package redis

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Address        string
	Password       string
	DB             int
	DefaultChannel string
}

type Client struct {
	client         *redis.Client
	defaultChannel string
}

func NewClient(opts Config) *Client {
	client := redis.NewClient(&redis.Options{
		Addr:     opts.Address,
		Password: opts.Password,
		DB:       opts.DB,
	})

	return &Client{
		client:         client,
		defaultChannel: opts.DefaultChannel,
	}
}

func (r *Client) Base() *redis.Client {
	return r.client
}

func (r *Client) Get(ctx context.Context, key string) ([]byte, error) {
	return r.client.Get(ctx, key).Bytes()
}

func (r *Client) Publish(ctx context.Context, message interface{}) error {
	_, err := r.client.Publish(ctx, r.defaultChannel, message).Result()
	return err
}

func (r *Client) Publishx(ctx context.Context, channel string, message interface{}) error {
	_, err := r.client.Publish(ctx, channel, message).Result()
	return err
}

// Get and unmarshall in T type
func Getx[T any](ctx context.Context, client *Client, key string, out *T) error {
	bytes, err := client.Get(ctx, key)
	if err != nil {
		return err
	}
	if out == nil {
		return errors.New("target cannot be nil")
	}
	if err := json.Unmarshal(bytes, out); err != nil {
		return err
	}
	return nil
}
