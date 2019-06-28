package selects

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compare two server and return the url of the fastest server", func(t *testing.T) {
		slowerServer := makeDelayedServer(10 * time.Millisecond)
		fasterServer := makeDelayedServer(0 * time.Millisecond)

		defer slowerServer.Close()
		defer fasterServer.Close()

		slowURL := slowerServer.URL
		fastURL := fasterServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}

		assertNoError(t, err)
	})

	t.Run("returns an error if a server doesn't respondse within 10s", func(t *testing.T) {
		timeout := 10 * time.Millisecond
		server := makeDelayedServer(timeout)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, timeout)

		assertError(t, err)
	})
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("expect no error, but get %v", err)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Errorf("expect error, but get nothing")
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
