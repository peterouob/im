package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB  *gorm.DB
	Red *redis.Client
	err error
)

func InitConfig() {
	//讀取配置文件
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app init ...")
}

func InitMysql() {
	//自定義日誌，打硬sql語句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢sql閥值
			LogLevel:      logger.Info, //級別
			Colorful:      true,        //彩色
		},
	)
	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})
	if err != nil {
		fmt.Println("init mysql Error", err)
	}
	fmt.Println("config init mysql ...")
}

func InitRedis() {
	Red = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.db"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	pong, err := Red.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("init redis Error", err)
	} else {
		fmt.Println("config init redis ...", pong)
	}
}

//Redis 訂閱和發布->提高性能
const (
	PublishKey = "websocket"
)

func Publish(ctx context.Context, channel, msg string) error {
	err = Red.Publish(ctx, channel, msg).Err()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Publish !", msg)
	return err
}

func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Red.Subscribe(ctx, channel)
	fmt.Println("Subscribe ...", ctx)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Subscribe !", msg.Payload)
	return msg.Payload, err
}
