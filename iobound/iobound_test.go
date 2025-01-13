package iobound

import "testing"

var urls = []string{
	"http://www.example.com",
	"http://example.org",
	"http://www.example.net",
}

func BenchmarkGetURLSequential(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		getURLSequential(urls)
	}
	b.StopTimer()
	b.ReportMetric(b.Elapsed().Seconds() / float64(b.N), "s/op") 
}

func BenchmarkGetURLConcurrent(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		getURLConcurrent(urls)
	}
	b.StopTimer()
	b.ReportMetric(b.Elapsed().Seconds() / float64(b.N), "s/op") 
}