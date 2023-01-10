package service

import "api/internal/library/libqueue"

type IQueue interface {
	libqueue.IQueuedJob
}

type sQueue struct {
	*libqueue.GQueue
}

var (
	cQ = sQueue{}
)

func Queue() IQueue {
	var ch = cQ
	ch.GQueue = libqueue.New()
	return &ch
}
