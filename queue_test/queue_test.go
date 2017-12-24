package queue_test

import (
	"testing"

	"github.com/Jrank2013/queue"
)

func TestLen(t *testing.T) {
	q := queue.NewQueue()
	q.AddItem("String", "Hello World", "Cool")

	item := q.Getitem()
	item.Remove()

	if q.Len() == 3 {
		t.Error("Did not remove an items")
	}

}
func TestNewQueue(t *testing.T) {
	q := queue.NewQueue()

	if q.Len() != 0 {
		t.Error("Queue was not initilzed properly")

	}
	testing.Coverage()

}

func TestGetitem(t *testing.T) {
	q := queue.NewQueue()

	q.AddItem("String")
	q.Getitem()
	item := q.Getitem()
	if item != nil {
		t.Error("Getitem returned item that wasn't visible")
	}
	q.AddItem("String")
	item = q.Getitem()
	if item.GetString() != "String" {
		t.Error("Returned non-exisitant item")
	}

}

func TestAddItem(t *testing.T) {
	q := queue.NewQueue()
	q.AddItem("String", "Hello World", "Cool")
	arr := []string{"String", "Hello World", "Cool"}

	if q.Len() != 3 {
		t.Error("Error inserting varadic number of args")
	}

	for i := 0; i < q.Len()-1; i++ {
		if q.Getitem().GetString() != arr[i] {
			t.Error("Getting wrong items from queue")
		}

	}
}

func TestRemove(t *testing.T) {
	q := queue.NewQueue()
	q.AddItem("String", "Hello World", "Cool")

	item := q.Getitem()
	item.Remove()

	if q.Len() == 3 {
		t.Error("Did not remove an items")
	}
	if q.Len() != 2 {
		t.Error("Removed wrong number of items")
	}

}
