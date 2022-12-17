package worker_store

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
)

/*
Worker Store Queue Definition
*/

type QueueModeEnum int64

const (
	Fifo QueueModeEnum = 0
	Filo               = 1
)

type WorkerStoreQueue struct {
	QueueId   int           `json:"queueId"`
	QueueName string        `json:"queueName"`
	Jobs      *list.List    `json:"jobs"`
	QueueMode QueueModeEnum `json:"queueMode"`
	Mutex     *sync.RWMutex
}

func (workerStoreQueue *WorkerStoreQueue) Get() (WorkerStoreJob, error) {
	var workerStoreJob WorkerStoreJob
	var getError error
	var element *list.Element
	if workerStoreQueue.Jobs.Len() > 0 {
		if workerStoreQueue.QueueMode == Fifo {
			element = workerStoreQueue.Jobs.Front()
		} else if workerStoreQueue.QueueMode == Filo {
			element = workerStoreQueue.Jobs.Back()
		} else {
			getError = errors.New(fmt.Sprintf("invalid queue mode %d", workerStoreQueue.QueueMode))
		}
	} else {
		getError = errors.New("no jobs available in queue")
	}

	if element != nil {
		workerStoreQueue.Jobs.Remove(element)
		workerStoreJob = element.Value.(WorkerStoreJob)
	}

	return workerStoreJob, getError
}
