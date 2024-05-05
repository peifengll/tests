# 1. lpush

这个功能用的时候要注意不能直接放结构体变量进去，
一般来说还是自己做解析，（记得可以自己为其实现 解析和读取的方法，然后就可以放了）

# 2. Exists方法

用于判断该键是否存在

```
func TestExistsValue(t *testing.T) {
	r := NewRedis()
	ctx := context.Background()
	result, err := r.Exists(ctx, "666").Result()
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(result)

}
```

存在会返回1，不存在返回0

# 3. HMGET

```go
result, err := r.HMGet(ctx, "666", "name", "age").Result()
```

666是键，是map的名字，后边的是这个map里边的键的名字
hmget中的m应该指的是many

对于hget，便是只取一个filed的值

hset，也是只设置一个



# 4. LRange

不存在这个键也不会报错

0，-1 取所有

返回的时候给的是字符串
