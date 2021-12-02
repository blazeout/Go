package main

import
(
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
		PoolSize: 100,
	})
	ctx , cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = redisdb.Ping(ctx).Result()
	return err
}

func  main() {
	ctx := context.Background()
	err := initRedis()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("链接redis成功")
	// zset
	key := "language_rank"
	languages := []*redis.Z {
		&redis.Z{Score: 99, Member: "Golang"},
		&redis.Z{Score: 98, Member: "Java"},
		&redis.Z{Score: 95, Member: "Python"},
		&redis.Z{Score: 90, Member: "C++"},
		&redis.Z{Score: 96, Member: "Rust"},
	}
	//ZADD
	redisdb.ZAdd(ctx,key,languages...)
	newScore, err := redisdb.ZIncrBy(ctx,key,1,"Golang").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("golang s score is %f now\n\n", newScore)
}
