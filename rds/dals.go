package rds

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rimvydaszilinskas/announcer-backend/models"
)

var ctx = context.Background()

func (r *RedisClient) StoreDeviceState(system *models.FullSystemDetails, device *models.Device) error {
	data, err := json.Marshal(system)

	if err != nil {
		return fmt.Errorf("error marshalling full system details - %s", err)
	}

	return r.redis.Set(ctx, device.GetRedisKey(), string(data), time.Minute*10).Err()
}

func (r *RedisClient) RetrieveDeviceState(device *models.Device) (*models.FullSystemDetails, error) {
	var details *models.FullSystemDetails
	key := device.GetRedisKey()
	value, err := r.redis.Get(ctx, key).Result()

	if err != nil {
		return nil, fmt.Errorf("error retrieving device with key %s - %s", key, err)
	}

	if err := json.Unmarshal([]byte(value), &details); err != nil {
		return nil, fmt.Errorf("error unmuarshalling data from redis")
	}

	return details, nil
}
