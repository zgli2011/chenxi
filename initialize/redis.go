package initialize

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

const (
	RedisMark               = "transaction_orchestration"
	NewJobQueueName         = RedisMark + ":NewJobQueue"
	SkipErrorJobQueueName   = RedisMark + ":SkipErrorJobQueue"
	ReTryJobQueueName       = RedisMark + ":ReTryJobQueue"
	ReDoJobQueueName        = RedisMark + ":ReDoJobQueue"
	ReTryJobNextQueueName   = RedisMark + ":ReTryJobNextQueue"
	ReDoJobNextQueueName    = RedisMark + ":ReDoJobNextQueue"
	ReTryJobPluginQueueName = RedisMark + ":ReTryJobPluginQueue"

	// NewJobRecordQueueName           = RedisMark + ":NewJobRecordQueue"
	// SkipErrorJobRecordQueueName     = RedisMark + ":SkipErrorJobRecordQueue"
	// ReTryJobRecordQueueName         = RedisMark + ":ReTryJobRecordQueue"
	// ReDoJobRecordQueueName          = RedisMark + ":ReDoJobRecordQueue"
	// ReTryJobRecordIPQueueName       = RedisMark + ":ReTryJobRecordIPQueue"
	// ReDoJobRecordIPQueueName        = RedisMark + ":ReDoJobRecordIPQueue"
	// ReTryJobRecordIPPluginQueueName = RedisMark + ":ReTryJobRecordIPPluginQueue"

	// NewJobStepRecordQueueName         = RedisMark + ":NewJobStepRecordQueue"
	// SkipErrorJobStepRecordQueueName   = RedisMark + ":SkipErrorJobStepRecordQueue"
	// ReTryJobStepRecordQueueName       = RedisMark + ":ReTryJobStepRecordQueue"
	// ReDoJobStepRecordQueueName        = RedisMark + ":ReDoJobStepRecordQueue"
	// ReTryJobStepRecordStepQueueName   = RedisMark + ":ReTryJobStepRecordStepQueue"
	// ReDoJobRecordStepQueueName        = RedisMark + ":ReDoJobRecordStepQueue"
	// ReTryJobStepRecordPluginQueueName = RedisMark + ":ReTryJobStepRecordPluginQueue"
)

type RedisOptoins struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	DB       int    `yaml:"db" json:"db"`
	PoolSize int    `yaml:"pool_size" json:"pool_size"`
	MinIdle  int    `yaml:"min_idle" json:"min_idle"`
}

func NewRedisClient() *redis.Client {
	redis_config := Config.Redis
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", redis_config.Host, redis_config.Port),
		PoolSize:     redis_config.PoolSize,
		MinIdleConns: redis_config.MinIdle,
		DB:           redis_config.DB,
		PoolTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,

		IdleCheckFrequency: 60 * time.Second,  //???????????????????????????????????????1?????????-1??????????????????????????????????????????????????????????????????????????????????????????
		IdleTimeout:        5 * time.Minute,   //?????????????????????5?????????-1??????????????????????????????
		MaxConnAge:         600 * time.Second, //??????????????????????????????????????????????????????????????????????????????????????????0??????????????????????????????????????????

		//????????????????????????????????????
		MaxRetries:      1,                      // ?????????????????????????????????????????????????????????0????????????
		MinRetryBackoff: 8 * time.Millisecond,   //????????????????????????????????????????????????8?????????-1??????????????????
		MaxRetryBackoff: 512 * time.Millisecond, //????????????????????????????????????????????????512?????????-1??????????????????
	})
	pong, err := RedisClient.Ping().Result()
	if err != nil {
		fmt.Println("redis ???????????????", pong, err)
		log.Panicf("redis connect failed: %v", err)
	}
	return RedisClient
}
