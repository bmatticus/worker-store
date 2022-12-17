package worker_store

type StoreType int64

const (
	Memory StoreType = 0
	Etcd             = 1
	Redis            = 2
)

type WorkerStoreType interface {
	AddQueue(WorkerStoreQueue) error
	GetQueue(string) (*WorkerStoreQueue, error)
}

func InitWorkerStore(storeType StoreType) WorkerStoreType {
	var workerStoreType WorkerStoreType
	if storeType == Memory {
		workerStoreType = InitMemoryWorkerStore()
	}

	return workerStoreType
}
