package main

import (
    "fmt"
    "time"
    "net/http"
    "github.com/bj-wangjia/mtggokit/metrics"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    go count()
    http.Handle("/metrics", promhttp.Handler())
    http.ListenAndServe("127.0.0.1:8888", nil)
}

func count() {
    Dimension := []string{"method", "code", "msg"}
    counter := metrics.NewCounterFromConfig("./demo_counter.yaml", Dimension)
    labels := prometheus.Labels{}
    labels["method"] = "Get"
    labels["code"] = "200"
    labels["msg"] = "succ"
    counter.With(labels).Add(1)
    labels["method"] = "Post"
    counter.With(labels).Add(2)
    time.Sleep(1000*time.Second)
}

