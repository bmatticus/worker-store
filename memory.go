package worker_store

/*
Memory store driver workers
*/

import (
	"errors"
	"fmt"
)

type WorkerStoreMemory struct {
	Queues map[string]*WorkerStoreQueue `json:"queues"`
}

func (workerStore *WorkerStoreMemory) AddQueue(queue WorkerStoreQueue) error {
	if _, ok := workerStore.Queues[queue.QueueName]; !ok {
		workerStore.Queues[queue.QueueName] = &queue
		return nil
	} else {
		return errors.New(fmt.Sprintf("queue withe name %s already exists", queue.QueueName))
	}
}

func (workerStore *WorkerStoreMemory) GetQueue(queueName string) (WorkerStoreQueue, error) {
	if queue, ok := workerStore.Queues[queueName]; ok {
		return queue, nil
	}
	return WorkerStoreQueue{}, errors.New(fmt.Sprintf("no queue found with name %s", queueName))
}

func InitMemoryWorkerStore() WorkerStoreType {
	return &WorkerStoreMemory{}
}
