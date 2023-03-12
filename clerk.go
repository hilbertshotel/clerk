package clerk

import (
	"net/http"
	"sync"
	"time"
)

// Holds necessary data for load testing
type Clerk struct {
	Request     *http.Request // request struct
	NumUsers    int           // number of users to make requests
	NumRequests int           // number of requests per user
	WaitTime    time.Duration // time to wait in between request
}

// Instantiate new Clerk struct with specified request and default fields
func New(req *http.Request) *Clerk {
	return &Clerk{
		Request:     req,
		NumUsers:    1,
		NumRequests: 1,
		WaitTime:    time.Second * 1,
	}
}

// Perform load testing on the request specified in Clerk
func (clerk *Clerk) Run() *Results {
	start := time.Now()
	var wg sync.WaitGroup
	var results Results

	for i := 0; i < clerk.NumUsers; i++ {
		wg.Add(1)

		go func(wg *sync.WaitGroup, pid int) {
			defer wg.Done()
			res := newResult(pid)

			n := clerk.NumRequests
			for n > 0 {
				t, err := RoundTrip(clerk.Request)
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
	results.RunTime = time.Since(start)

	return &results
}

// Times an http request
func RoundTrip(req *http.Request) (time.Duration, error) {
	start := time.Now()
	_, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return time.Duration(0), err
	}

	return time.Since(start), nil
}
