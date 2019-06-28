package selects

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecTimeout = 10 * time.Second

// Racer race to website, and return the fastest to response
// it will return error if it takes longer than 10s
func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecTimeout)
}

// ConfigurableRacer a Racer with timeout that can be configured
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	// listento multiple channel
	// return the first one to finished
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("time out waiting for '%s' and '%s'", a, b)
	}
}

// wait till get http get response,
// then send the signal through channel
func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()

	return ch
}
