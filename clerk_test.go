package clerk

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func Test(t *testing.T) {

	req, err := http.NewRequest("GET", "https://coingecko.com", nil)
	if err != nil {
		panic(err)
	}

	clerk := New(req)
	clerk.NumUsers = 3
	clerk.NumRequests = 2
	clerk.WaitTime = time.Millisecond * 500

	results := clerk.Run()

	fmt.Println(results.RunTime)
	for _, res := range results.List {
		fmt.Println(res)
	}
}
