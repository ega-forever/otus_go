package queue

type Iterator struct {
	current *QueueItem
	queue   *Queue
}

func (iterator *Iterator) HasNext() bool {
	return iterator.current != nil && iterator.current.nextItem != nil
}

func (iterator *Iterator) GetNext() *QueueItem {

	if !iterator.HasNext() {
		return nil
	}

	iterator.current = iterator.current.nextItem

	return iterator.current
}

func (iterator *Iterator) GetCurrent() *QueueItem {
	return iterator.current
}

func (iterator *Iterator) HasPrev() bool {
	return iterator.current != nil && iterator.current.prevItem != nil
}

func (iterator *Iterator) GetPrev() *QueueItem {

	if !iterator.HasPrev() {
		return nil
	}

	iterator.current = iterator.current.prevItem

	return iterator.current
}
