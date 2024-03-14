// monitoring.go
package monitoring

import (
    "log"
    "net/http"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var counter = prometheus.NewCounter(prometheus.CounterOpts{
    Name: "requests_total",
    Help: "Total number of requests.",
})

func IncrementCounter() {
    counter.Inc()
}

func SetupMetrics() {
    prometheus.MustRegister(counter)

    
    

    http.Handle("/metrics", promhttp.Handler())

    log.Println("Monitoring setup complete")
}
