Clerk is a simple load testing package

How to use:
```go
package main

import (
    "github.com/hilbertshotel/clerk"
)

func main() {
    // Instantiate new Clerk struct with the request url as argument
    clerk := clerk.New("https://coingecko.com")

    // Modify parameters
    clerk.NumUsers = 200                    // default=1
    clerk.NumRequests = 10                  // default=1
    clerk.WaitTime = time.Millisecond * 500 // default=1s

    // Run load testing
    results := clerk.Run()

    // Check results slice for response times & errors
    for _, res := range results.List {
        fmt.Println(res)
    } 
}
```

Type annotations:
```go
type Clerk struct {
	URL         string        // request url
	NumUsers    int           // number of users to make requests
	NumRequests int           // number of requests per user
	WaitTime    time.Duration // time to wait in between request
}

type Result struct {
	Pid       int
	RespTimes []time.Duration
	Errors    []error
}

type Results struct {
	List  []Result
	mutex sync.Mutex
}
```