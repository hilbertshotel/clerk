package clerk

import (
	"sync"
	"time"
)

// Holds result data for each user/goroutine
type Result struct {
	Pid       int
	RespTimes []int64
	Errors    []error
}

// Holds list of results and a lock for concurrent writing
type Results struct {
	List    []Result
	RunTime time.Duration
	mutex   sync.Mutex
}

// Instantiate new Result struct with specified process id
func newResult(pid int) Result {
	return Result{
		Pid: pid,
	}
}

// Concurrently safe method for adding a result to results list
func (res *Results) add(result Result) {
	res.mutex.Lock()
	res.List = append(res.List, result)
	res.mutex.Unlock()
}
