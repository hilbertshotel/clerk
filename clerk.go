package clerk

import (
	"net/http"
	"sync"
	"time"
)

// Holds necessary data for load testing
type Clerk struct {
	URL         string        // request url
	NumUsers    int           // number of users to make requests
	NumRequests int           // number of requests per user
	WaitTime    time.Duration // time to wait in between request
}

// Instantiate new Clerk struct with specified url and default fields
func New(url string) *Clerk {
	return &Clerk{
		URL:         url,
		NumUsers:    1,
		NumRequests: 1,
		WaitTime:    time.Second * 1,
	}
}

// Perform load testing on the url specified in Clerk
func (clerk *Clerk) Run() *Results {
	var wg sync.WaitGroup
	var results Results

	for i := 0; i < clerk.NumUsers; i++ {
		wg.Add(1)

		go func(wg *sync.WaitGroup, pid int) {
			defer wg.Done()
			res := newResult(pid)

			n := clerk.NumRequests
			for n > 0 {
				t, err := RoundTrip(clerk.URL)
				if err != nil {
					res.Errors = append(res.Errors, err)
					continue
				}
				res.RespTimes = append(res.RespTimes, t)

				time.Sleep(clerk.WaitTime)
				n--
			}

			results.add(res)
		}(&wg, i+1)

	}

	wg.Wait()

	return &results
}

// Send and time a get request to the specified url
func RoundTrip(url string) (time.Duration, error) {
	var t time.Duration

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return t, err
	}

	start := time.Now()
	_, err = http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return t, err
	}

	return time.Since(start), nil
}
