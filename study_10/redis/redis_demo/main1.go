package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6399",
		Password: "", // 密码
		DB:       0,  // 使用的库
	})
	_, err = redisdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func setStr() {
	err := redisdb.Set("test", 100, 0).Err()
	if err != nil {
		fmt.Println("redis set err")
		return
	}
	fmt.Println("设置成功")
}

func getStr() {
	value, err := redisdb.Get("test").Result()
	if err != nil {
		fmt.Println("redis set err")
		return
	}
	fmt.Println("获取成功:", value)
}

func zsetDome() {
	// 初始化一个ke
	zsetkey := "language_rank"
	languages := []*redis.Z{
		&redis.Z{Score: 92, Member: "php"},
		&redis.Z{Score: 94, Member: "golang"},
		&redis.Z{Score: 97, Member: "java"},
		&redis.Z{Score: 99, Member: "c"},
	}
	n, err := redisdb.ZAdd(zsetkey, languages...).Result()
	if err != nil {
		fmt.Println("redis ZAdd ", err)
		return
	}
	fmt.Printf("redis key %s , 新增了 %d 元素\n", zsetkey, n)
	// 给某个元素加个分数
	ret, err := redisdb.ZIncrBy(zsetkey, 3, "php").Result()
	if err != nil {
		fmt.Println("ZIncrBy err", err)
		return
	}
	fmt.Printf("修改分数成功，当前分数：%v\n", ret)
	// 取分数最高的3个
	ret1, err := redisdb.ZRevRangeWithScores(zsetkey, 0, 2).Result()
	if err != nil {
		fmt.Println("ZRangeByScoreWithScores", ret1)
		return
	}
	fmt.Println("分数最高的3位")
	for _, v := range ret1 {
		fmt.Println(v.Member, v.Score)
	}
	// 获取分数在95-100的
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	fmt.Println("获取分数在95-100的")
	ret2, err := redisdb.ZRangeByScoreWithScores(zsetkey, op).Result()
	for _, v := range ret2 {
		fmt.Println(v.Member, v.Score)
	}
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Println("实例化redis有误", err)
	}
	fmt.Println("链接redis成功")
	// setStr()
	// getStr()
	zsetDome()

}
