package libqueue

type IQueuedJob interface {
	Push(data interface{})
	Consumption()
}
