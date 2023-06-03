package concurrent

import (
	"math/rand"
	"sync"
)

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func balancingRoutine(e *execService, queueId int, thresholdBalance int, capapcity int, wg *sync.WaitGroup) {
	queue := e.localQueues[queueId]

	for !(e.done && e.globalQueue.IsEmpty() && queue.IsEmpty()) {
		// local queues push tasks from local queue into their own local queues
		for i := 0; i < e.threshold; i++ {
			elem := e.globalQueue.PopTop()
			if elem == nil {
				break
			}
			queue.PushBottom(elem)
		}

		// loadbalancing algorithm for parallel synchronization
		// If randSize chosen is equal to the queue size (probability = 1/size)
		// If randSize == size then a random victim is chosen
		// If the diff between victim queue size and curr queue size is greater than
		// or equal to threshold balance then balancing occurs between the two queues
		loadBalancing := func(e *execService, currId int) {
			curr_queue := e.localQueues[queueId]
			size := curr_queue.GetSize()

			randSize := rand.Intn(size + 1)

			if randSize == size {
				// choose a random victim
				victim := rand.Intn(capapcity)
				victimSize := e.localQueues[victim].GetSize()

				balanceQueues := func(id_max, id_min, diff int) {
					balaceAmount := diff / 2
					for i := 0; i < balaceAmount; i++ {
						e.localQueues[id_min].PushBottom(
							e.localQueues[id_max].PopTop(),
						)
					}
				}

				diff := absDiffInt(size, victimSize)

				if diff >= thresholdBalance {
					if size > victimSize {
						balanceQueues(queueId, victim, diff)
					} else {
						balanceQueues(victim, queueId, diff)

					}
				}
			}
		}

		for !queue.IsEmpty() {
			// local queue process its own tasks
			elem := queue.PopBottom()

			if elem == nil {
				break
			}

			currTask, currChan := elem.(future).task, elem.(future).c

			if obj, ok := currTask.(interface{ Call() interface{} }); ok {
				val := obj.Call()
				currChan <- val
			}

			if obj, ok := currTask.(interface{ Run() }); ok {
				obj.Run()
				currChan <- nil
			}

		}

		loadBalancing(e, queueId)
	}
	wg.Done()
}

// NewWorkBalancingExecutor returns an ExecutorService that is implemented using the work-balancing algorithm.
// @param capacity - The number of goroutines in the pool
// @param threshold - The number of items that a goroutine in the pool can
// grab from the executor in one time period. For example, if threshold = 10
// this means that a goroutine can grab 10 items from the executor all at
// once to place into their local queue before grabbing more items. It's
// not required that you use this parameter in your implementation.
// @param thresholdBalance - The threshold used to know when to perform
// balancing. Remember, if two local queues are to be balanced the
// difference in the sizes of the queues must be greater than or equal to
// thresholdBalance. You must use this parameter in your implementation.
func NewWorkBalancingExecutor(capacity, threshold, thresholdBalance int) ExecutorService {
	/** TODO: Remove the return nil and implement this function **/

	// create num of localqueues equal to the capacity length
	var lst []DEQueue = make([]DEQueue, capacity)

	for i := 0; i < capacity; i++ {
		lst[i] = NewUnBoundedDEQueue()
	}

	e := &execService{
		capapcity:   capacity,
		threshold:   threshold,
		globalQueue: NewUnBoundedDEQueue(),
		localQueues: lst,
		done:        false,
		wg:          &sync.WaitGroup{},
	}

	var wg sync.WaitGroup

	e.wg.Add(1)

	go func() {
		for i := 0; i < capacity; i++ {
			wg.Add(1)
			go balancingRoutine(e, i, thresholdBalance, capacity, &wg)
		}
		wg.Wait()
		e.wg.Done()
	}()

	return e
}
