package concurency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "https://unreachable.com"
}

func TestCheckWebSites(t *testing.T) {
	urlsToCheck := []string{"https://www.miniamatch.com", "https://www.google.com", "https://unreachable.com"}
	want := map[string]bool{"https://www.miniamatch.com": true, "https://www.google.com": true, "https://unreachable.com": false}
	got := CheckWebSites(mockWebsiteChecker, urlsToCheck)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("got %v want %v", got, want)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebSites(slowStubWebsiteChecker, urls)
	}
}
