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

		IdleCheckFrequency: 60 * time.Second,  //闲置连接检查的周期，默认为1分钟，-1表示不做周期性检查，只在客户端获取连接时对闲置连接进行处理。
		IdleTimeout:        5 * time.Minute,   //闲置超时，默认5分钟，-1表示取消闲置超时检查
		MaxConnAge:         600 * time.Second, //连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接

		//命令执行失败时的重试策略
		MaxRetries:      1,                      // 命令执行失败时，最多重试多少次，默认为0即不重试
		MinRetryBackoff: 8 * time.Millisecond,   //每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
		MaxRetryBackoff: 512 * time.Millisecond, //每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔
	})
	pong, err := RedisClient.Ping().Result()
	if err != nil {
		fmt.Println("redis 连接失败：", pong, err)
		log.Panicf("redis connect failed: %v", err)
	}
	return RedisClient
}
