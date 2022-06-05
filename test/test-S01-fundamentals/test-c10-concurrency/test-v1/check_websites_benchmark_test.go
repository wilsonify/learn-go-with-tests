package concurrency

import (
	"testing"
	"time"

	concurrency "learn.go/S01-fundamentals/c10-concurrency/v1"
)

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		concurrency.CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
