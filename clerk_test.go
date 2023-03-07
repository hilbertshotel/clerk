package clerk

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	clerk := New("https://coingecko.com")
	clerk.NumUsers = 200
	clerk.NumRequests = 1
	clerk.WaitTime = time.Millisecond * 500

	results := clerk.Run()

	for _, res := range results.List {
		fmt.Println(res)
	}
}
