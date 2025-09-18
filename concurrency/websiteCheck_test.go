package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://fun.org"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.wyx.com",
		"waat://fun.org",
	}

	actualResults := CheckWebsites(mockWebsiteChecker, websites)

	want := len(websites)
	got := len(actualResults)
	if want != got {
		t.Fatalf("wanted %v, got %v", want, got)
	}

	expectedResults := map[string]bool{
		"http://google.com":   true,
		"http://blog.wyx.com": true,
		"waat://fun.org":      false,
	}

	if !reflect.DeepEqual(actualResults, expectedResults) {
		t.Fatalf("wanted %v, got %v", expectedResults, actualResults)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	url := make([]string, 100)
	for i := 0; i < len(url); i++ {
		url[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, url)
	}
}
