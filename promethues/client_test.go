package promethues

import "testing"

func TestPromClient(t *testing.T) {
	PromRun()
}
func TestApiClient(t *testing.T) {
	p := NewApiClient()
	p.Series("up", "node_cpu")
}
func TestPromeClient_Query(t *testing.T) {
	p := NewApiClient()
	p.Query(`http_requests_total{code="200"}`)
}
