package racing

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

    slowServer := makeServer(10 * time.Millisecond)
    fastServer := makeServer(0 * time.Millisecond)
    defer slowServer.Close()
    defer fastServer.Close()

    // slowUrl := "http://www.github.com"
    // fastUrl := "http://www.google.com"

    slowUrl := slowServer.URL
    fastUrl := fastServer.URL

    expected := fastUrl
    got := WebRacer(slowUrl, fastUrl)

    if expected != got {
        t.Errorf("expected %q but got %q", expected, got )
    }
}

func makeServer(delay time.Duration) *httptest.Server {
    return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(delay)
        w.WriteHeader(http.StatusOK)
    }))
}