package libqueue

import (
	"fmt"
	"sync"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/container/gqueue"
)

type GQueue struct {
	queue *gqueue.Queue
}

var (
	queue *gqueue.Queue
	once  sync.Once
)

func New() *GQueue {
	once.Do(func() {
		queue = gqueue.New()
	})
	return &GQueue{
		queue: queue,
	}
}

func (q *GQueue) Push(value interface{}) {
	q.queue.Push(value)
}

func (q *GQueue) Consumption() {

	fmt.Println("监听队列")
	for {
		if v := q.queue.Pop(); v != nil {
			g.Dump(v)
		}
	}
}
