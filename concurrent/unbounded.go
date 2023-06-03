package concurrent

import "sync"

/**** YOU CANNOT MODIFY ANY OF THE FOLLOWING INTERFACES/TYPES ********/
type Task interface{}

type DEQueue interface {
	PushBottom(task Task)
	IsEmpty() bool //returns whether the queue is empty
	PopTop() Task
	PopBottom() Task
	GetSize() int
}

/******** DO NOT MODIFY ANY OF THE ABOVE INTERFACES/TYPES *********************/

type UnBoundedDEQueue struct {
	head *Node
	tail *Node
	size int
	mt   *sync.Mutex
}

type Node struct {
	task *Task
	next *Node
	prev *Node
}

// NewUnBoundedDEQueue returns an empty UnBoundedDEQueue
func NewUnBoundedDEQueue() DEQueue {
	/** TODO: Remove the return nil and implement this function **/
	return &UnBoundedDEQueue{nil, nil, 0, &sync.Mutex{}}
}

func (q *UnBoundedDEQueue) GetSize() int {
	return q.size
}
func (q *UnBoundedDEQueue) PushBottom(task Task) {
	q.mt.Lock()
	defer q.mt.Unlock()

	n := &Node{task: &task}
	if q.head == nil { // queue is empty
		q.head = n
	} else {
		q.tail.next = n //	Add at the end
		n.prev = q.tail // Update the prev of last node
	}
	q.tail = n
	q.size++
}

func (q *UnBoundedDEQueue) PopTop() Task {
	q.mt.Lock()
	defer q.mt.Unlock()

	if q.head == nil { // empty queue
		return nil
	} else if q.head == q.tail {
		v := q.head

		q.head, q.tail = nil, nil
		q.size--

		return *v.task
	} else {
		v := q.head

		prev := q.head.prev
		next := q.head.next

		if next != nil { // only element
			next.prev = prev
		}

		q.head = next
		q.size--

		return *v.task
	}

}

func (q *UnBoundedDEQueue) PopBottom() Task {
	q.mt.Lock()
	defer q.mt.Unlock()

	if q.head == nil {
		return nil
	} else if q.head == q.tail {
		v := q.tail
		q.head, q.tail = nil, nil
		q.size--
		return *v.task
	} else {
		v := q.tail
		prev := q.tail.prev
		next := q.tail.next
		if prev != nil {
			prev.next = next
		}
		q.tail = prev
		q.size--
		return *v.task
	}
}

func (q *UnBoundedDEQueue) IsEmpty() bool {
	q.mt.Lock()
	defer q.mt.Unlock()
	return q.size == 0
}
