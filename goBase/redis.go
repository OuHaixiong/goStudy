package main

import (
	"fmt"
	"gopkg.in/redis.v5"
	"github.com/jie123108/glog"
	// "time"
)

func main() {
	// client, err := redis.Dial("tcp", "172.17.6.140:6379") // 这种方式是在包："github.com/garyburd/redigo/redis"
	// if err != nil {
	// 	fmt.Println("Connect to Redis error:", err)
	// 	return
	// }
	// defer client.Close()

	client := redis.NewClient(&redis.Options{Addr:"172.17.6.140:6379", Password:"123456", DB:12}) // no password set ""  DB: 0, use default DB
	pong, err := client.Ping().Result()
	if err != nil {
        glog.Error("Redis connect is error:", err, " , pong:", pong)
		return
	}
	defer client.Close()
    fmt.Println("Is Connected")	

	// keyTest will expire in an hour.
	// err = client.Set("keyTest", "test Value", time.Hour).Err() // 设置一个string型数据，在一个小时后过期。 如果设置为0，则永不过期
	// if (err != nil) {
    //     panic(err)
	// }
	testValue := client.Get("keyTest") // 这里返回的是 *StringCmd
	glog.Error(testValue)

	val, err := client.Get("keyTest").Result()
	if err == redis.Nil {
        fmt.Println("key does not exists")
	} else if err != nil {
		panic(err)
	}
    glog.Error(val) // test Value

	key := "testkey"
	client.LSet(key, 0, "XXXXX") // 这里返回：*StatusCmd  .   给list中index位置的元素赋值
	value := client.LIndex(key, 0).Val(); // LIndex(key string, index int64) *StringCmd 返回名称为key的list中index位置的元素
	// 特别注意上面的了，如果index不存在是不会添加的
	// if err != nil {
	// 	panic(err)
	// }
	glog.Error(value)

    testList(client);

}

/**
list 设置
**/
func testList(client *redis.Client) {
	// err := client.LPush("test-list", 1, 2, 3, 4).Err() //可以支持多个push 。 lpush(key, value)：在名称为key的list头添加一个值为value的 元素
	// fmt.Println(err) // <nil>
	// value := client.LPop("test-list").Val() // 返回并删除list中的首元素
	// fmt.Println(value) // 4
	// statusCmd := client.LTrim("test-list", 0, -1)
	length, err := client.LLen("test-list").Result() // 返回list的长度
	if (err != nil) {
        panic(err)
	}
	fmt.Println(length)

	value, err := client.RPop("test-list").Result() // 返回并删除list中的尾元素
	if err != nil {
		panic(err)
	}
	fmt.Println(value)
	// lrange(key, start, end)：返回名称为key的list中start至end之间的元素

	value = client.LIndex("test-list", 0).Val(); // LIndex(key string, index int64) *StringCmd 返回名称为key的list中index位置的元素
	glog.Error(value)
	
}