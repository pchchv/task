// Provides a simple "data storage" in memory for tasks.
// Tasks are uniquely identified by numeric identifiers.

package taskstore

import (
	"sync"

	"github.com/pchchv/task/graph/model"
)

// Simple in-memory task database;
// methods are safe for simultaneous invocation.
type TaskStore struct {
	sync.Mutex

	tasks  map[int]*model.Task
	nextId int
}

func New() *TaskStore {
	ts := &TaskStore{}
	ts.tasks = make(map[int]*model.Task)
	ts.nextId = 0

	return ts
}
