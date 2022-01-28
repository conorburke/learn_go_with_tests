package sync

import "sync"

func main() {

}

type Counter struct {
	value int
	mutex sync.Mutex
}

func (c *Counter) Value() int {
	return c.value
}

func (c *Counter) Inc() {
    c.mutex.Lock()
    defer c.mutex.Unlock()
	c.value++
}

func NewCounter() *Counter {
    return &Counter{}
}
