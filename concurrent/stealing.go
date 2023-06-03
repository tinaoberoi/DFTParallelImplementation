package concurrent

import (
	"math/rand"
	"sync"
	"time"
)

func getRandomIntWithBlacklist(max int, blacklisted int) int {
	// loop until an n is generated that is not in the blacklist
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(max) // yields n such that min <= n <= max
		if n != blacklisted {
			// fmt.Println("n = ", n, "blacklisted = ", blacklisted)
			return n
		}
	}
}

// NewWorkStealingExecutor returns an ExecutorService that is implemented using the work-stealing algorithm.
// @param capacity - The number of goroutines in the pool
// @param threshold - The number of items that a goroutine in the pool can
// grab from the executor in one time period. For example, if threshold = 10
// this means that a goroutine can grab 10 items from the executor all at
// once to place into their local queue before grabbing more items. It's
// not required that you use this parameter in your implementation.
func workerRoutine(e *execService, queueId int, wg *sync.WaitGroup) {
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

		// If queue is empty that is no tasks could be taken from global queue
		// then stealing algorithm triggers, a random victim is chosen and
		// tasks = threshold are taken from the victim to process
		if queue.IsEmpty() {
			victimId := getRandomIntWithBlacklist(e.capapcity, queueId)
			vitcimQueue := e.localQueues[victimId]

			for j := 0; j < e.threshold; j++ {
				stolenElem := vitcimQueue.PopTop()
				if stolenElem == nil {
					break
				}
				queue.PushBottom(stolenElem)
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
	}
	wg.Done()
}

func NewWorkStealingExecutor(capacity, threshold int) ExecutorService {
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
			go workerRoutine(e, i, &wg)
		}
		wg.Wait()
		e.wg.Done()
	}()

	return e
}
