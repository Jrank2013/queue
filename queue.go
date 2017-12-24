package queue

import (
	"time"

	"github.com/rs/xid"
)

type queueItem struct {
	visibilityTimeout *time.Duration
	member            *string
	visibile          *bool
	UUID              *string
	parentQueue       *Queue
}

//Queue contains queueItems
//queueItems contain visibilityTimeout which sets how long before an item is Visibile
//before it can be accessed again
//member is the data in the queue
//visibile is weather or not the item is visible
type Queue struct {
	queue []queueItem
}

type access interface {
	Additem(toAdd string)
	remove()
	getitem() string
	SetVisibilityTimeout(milliseconds float64)
	getString() string
}

//AddItem adds items to end of queue
func (q *Queue) AddItem(items ...string) {
	for _, item := range items {
		q.queue = append(q.queue, queueItem{member: newString(item), visibilityTimeout: newTime(10),
			visibile: newBool(true), UUID: newUUID(), parentQueue: q})
	}
}

//Getitem Gets the next items from the queue that is visible
//If nothing in queue is visible returns nil
func (q *Queue) Getitem() *queueItem {
	for _, item := range q.queue {
		if !(*item.visibile) {
			continue
		}
		*item.visibile = false
		item.startTimer()
		return &item
	}
	return nil
}

//Removes removes item from queue
func (q *queueItem) Remove() {
	for i, item := range q.parentQueue.queue {
		if q.UUID == item.UUID {
			q.parentQueue.queue = append(q.parentQueue.queue[:i], q.parentQueue.queue[i+1:]...)
		}
	}
}

//Len returns length of queue
func (q *Queue) Len() int {
	return len(q.queue)
}

func (q *queueItem) SetVisibilityTimeout(seconds time.Duration) {
	q.visibilityTimeout = newTime(seconds)
}

func (q *queueItem) startTimer() {
	timer := time.NewTimer(*q.visibilityTimeout)
	go func() {
		<-timer.C
		*q.visibile = true
	}()
}

func newUUID() *string {
	guid := xid.New().String()
	return &guid
}

func (q *queueItem) GetString() string {
	return *q.member
}

//NewQueue returns Queue with no items
func NewQueue() Queue {
	q := make([]queueItem, 0)
	return Queue{queue: q}
}

func newBool(value bool) *bool {
	b := value
	return &b
}

func newTime(seconds time.Duration) *time.Duration {
	b := time.Second * seconds
	return &b
}

func newString(value string) *string {
	b := value
	return &b
}
