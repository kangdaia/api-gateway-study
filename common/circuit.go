package common

import (
	"time"

	"github.com/sony/gobreaker/v2"
)

// circuit breaker: 장애가 전파하는 것을 (실패하는 동작을 불필요햐게 시도하는 것을) 방지하기 위한 디자인 패텬
var CB *gobreaker.CircuitBreaker[[]byte]

func init() {
	var st gobreaker.Settings
	st.Name = "study_circuit_breaker"
	st.MaxRequests = 0
	st.Interval = time.Second

	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		ratio := float64(counts.TotalFailures) / float64(counts.TotalSuccesses)
		return counts.Requests >= 3 && ratio >= 0.6 // stop condition
	}

	CB = gobreaker.NewCircuitBreaker[[]byte](st)
}
