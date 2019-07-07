package queue

type Queue struct {
	head *QueueItem
	tail *QueueItem
}

func (queue *Queue) GetIterator() Iterator {
	iterator := Iterator{queue.tail, queue}
	return iterator
}

func (queue *Queue) Add(item int) {
	lastElem := queue.head
	newElem := QueueItem{lastElem, item, nil}

	if lastElem != nil {
		lastElem.nextItem = &newElem
		queue.head = &newElem
	} else {
		queue.tail = &newElem
		queue.head = &newElem
	}
}

func (queue *Queue) Get(index int) *QueueItem {

	counter := 0
	iterator := Iterator{queue.tail, queue}

	for iterator.HasNext() && counter != index {
		iterator.GetNext()
		counter++
	}

	if counter != index {
		return nil
	}

	return iterator.GetCurrent()
}

func (queue *Queue) Remove(index int) {

	counter := 0
	iterator := Iterator{queue.tail, queue}

	for iterator.HasNext() && counter != index {
		iterator.GetNext()
		counter++
	}

	current := iterator.GetCurrent()

	if current == nil {
		return
	}

	if current.prevItem == nil && current.nextItem == nil {
		queue.head = nil
		queue.tail = nil
		return
	}

	if current.prevItem == nil && current.nextItem != nil {
		queue.tail = current.nextItem
		queue.tail.prevItem = nil
		return
	}

	if current.prevItem != nil && current.nextItem == nil {
		current.prevItem.nextItem = nil
		queue.head = current.prevItem
		return
	}

	if current.prevItem != nil && current.nextItem != nil {
		current.prevItem.nextItem = current.nextItem
		current.nextItem.prevItem = current.prevItem
	}

}

func (queue *Queue) Len() int {

	if queue.head == nil {
		return 0
	}

	count := 1
	iterator := Iterator{queue.tail, queue}

	for iterator.HasNext() {
		iterator.GetNext()
		count++
	}

	return count
}
