package main

import (
	"context"
	"testing"
)

const setKey = "i_am_z_set"

func TestAdd(t *testing.T) {
	r := NewRedis()
	ctx := context.Background()
	r.SAdd(ctx, setKey, 99)
	r.SAdd(ctx, setKey, 98)
}

func TestRem(t *testing.T) {
	r := NewRedis()
	ctx := context.Background()
	// 移除目标元素
	r.SRem(ctx, setKey, 98)
}

/*
func GetRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

// redisSetTest Set集合(无序不重复)
func redisSetTest(cli *redis.Client) {
	// SAdd:添加（可一次添加多个）
	cli.SAdd("city", "西二旗", "东京", "纽约", "首尔", "新德里")

	// SRem:移除(可一次移除多个)
	cli.SRem("city", "东京", "首尔")

	//SMembers:查看set中的全部信息
	fmt.Println(cli.SMembers("city").Val()) // [新德里 纽约 西二旗]

	// 判断set中是否存在某个值
	fmt.Println(cli.SIsMember("city", "北京").Val()) //false

	//随机获取set集合中的一个元素
	fmt.Println(cli.SRandMember("city").Val()) //西二旗

	//随机获取指定个数的元素
	fmt.Println(cli.SRandMemberN("city", 2).Val()) //[新德里 西二旗]

	cli.SAdd("city", "莫斯科", "平襄", "雅加达", "马尼拉")
	fmt.Println(cli.SMembers("city").Val()) //[莫斯科 马尼拉 纽约 新德里 西二旗 平襄 雅加达]
	//随机删除一条数据
	cli.SPop("city")
	fmt.Println(cli.SMembers("city").Val()) //[莫斯科 纽约 马尼拉 平襄 雅加达 新德里]
	//随机删除指定条数数据
	cli.SPopN("city", 3)
	fmt.Println(cli.SMembers("city")) // [马尼拉 纽约 平襄]

	//将指定的值移动到指定集合(指定集合不存在则创建)
	b, err := cli.SMove("city", "language", "纽约").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(b)                              	//true
	fmt.Println(cli.SMembers("city").Val())     //[马尼拉 平襄]
	fmt.Println(cli.SMembers("language").Val()) //[纽约]

	//数字集合类
	cli.SAdd("set1","a","b","c","d")
	cli.SAdd("set2","a","c","e","f")

	//差集（注意以谁为基准,也就是第一个参数是谁！！！）
	fmt.Println(cli.SDiff("set1","set2").Val())	//[b d]
	fmt.Println(cli.SDiff("set2","set1").Val())	//[e f]
	//交集
	fmt.Println(cli.SInter("set1","set2").Val())//[a c]
	//并集
	fmt.Println(cli.SUnion("set1","set2").Val())//[c a b f d e]

}

func main() {
	rdb := GetRedisClient()
	defer rdb.Close()
	pong := rdb.Ping().Val()
	fmt.Printf("数据连接状态：%v\n", pong) // 数据连接状态：PONG

	redisSetTest(rdb)
}

*/
