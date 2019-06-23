package queue

type QueueItem struct {
	prevItem *QueueItem
	Value    int
	nextItem *QueueItem
}
