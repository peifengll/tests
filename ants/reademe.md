# ants 是什么

ants 是一个协程池的实现

# 为什么要有golang的协程池

Golang 提供的 go 关键字能很方便地将多个协程塞在一个进程中。但是在实际开发过程中，我们容易遇到协程滥用的问题。



# 一些参数
1. ExpiryDuration：过期时间。表示 goroutine 空闲多长时间之后会被ants池回收
2. PreAlloc：预分配。调用NewPool()/NewPoolWithFunc()之后预分配worker（管理一个工作 goroutine 的结构体）切片。而且使用预分配与否会直接影响池中管理worker的结构。
3. MaxBlockingTasks：最大阻塞任务数量。即池中 goroutine 数量已到池容量，且所有 goroutine 都处理繁忙状态，这时到来的任务会在阻塞列表等待。这个选项设置的是列表的最大长度。阻塞的任务数量达到这个值后，后续任务提交直接返回失败
4. Nonblocking：池是否阻塞，默认阻塞。提交任务时，如果ants池中 goroutine 已到上限且全部繁忙，阻塞的池会将任务添加的阻塞列表等待（当然受限于阻塞列表长度，见上一个选项）。非阻塞的池直接返回失败
5. PanicHandler：panic 处理。遇到 **panic 会调用这里设置的处理函数**
6. Logger：指定日志记录器

