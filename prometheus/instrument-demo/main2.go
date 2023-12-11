package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// 使用、暴露和创建一个非全局注册对象
func main() {
	registry := prometheus.NewRegistry()
	// 添加process和go运行时指标到自定义的注册表中
	//registry.MustRegister(prometheus.NewProcessCollector(
	//	prometheus.ProcessCollectorOpts{}))
	//registry.MustRegister(prometheus.NewGoCollector())
	//	 创建一个简单的gauge指标
	temp := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "home_temperature_celsius",
		Help: "The current temperature in degrees Celsius.",
	})
	registry.MustRegister(temp)
	temp.Set(39)

	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{
		Registry: registry,
	}))
	http.ListenAndServe(":8080", nil)

}
