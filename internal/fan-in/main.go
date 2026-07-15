package fan_in

import "sync"

func fanIn(chans ...<-chan int) <-chan int {
	res := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(chans))

	for _, ch := range chans {
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				res <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}
