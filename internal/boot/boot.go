package boot

import "api/internal/library/libqueue"

func init() {
	//fmt.Println("实时监听")
	go func() {
		libqueue.New().Consumption()
	}()
}
