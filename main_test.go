package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// tun go test to test this func
// func TestHandler(t *testing.T) {
// 	cases := []struct{ in, out string }{
// 		{"campoy@golang.org", "gopher campoy"},
// 		{"something", "dear something"},
// 	}

// 	for _, c := range cases {
// 		req, err := http.NewRequest(
// 			http.MethodGet,
// 			"http://localhost:80890/"+c.in,
// 			nil,
// 		)

// 		if err != nil {
// 			t.Fatalf("could not create request: %v", err)
// 		}

// 		rec := httptest.NewRecorder()
// 		handler(rec, req)

// 		if rec.Code != http.StatusOK {
// 			t.Errorf("expected status 200: got %d", rec.Code)
// 		}
// 		if strings.Contains(rec.Body.String(), c.out) {
// 			t.Errorf("unexpected body in response: %q", rec.Body.String())
// 		}
// 	}
// }

// create a Benchmark
// cmd: go test -bench .
func BenchmarkHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest(
			http.MethodGet,
			"http://localhost:8080/campoy@golang.org",
			nil,
		)

		if err != nil {
			b.Fatalf("could not create request: %v", err)
		}

		rec := httptest.NewRecorder()
		handler(rec, req)

		if rec.Code != http.StatusOK {
			b.Errorf("expected status 200: got %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), "gopher campoy") {
			b.Errorf("unexpected body in response: %q", rec.Body.String())
		}
	}
}
