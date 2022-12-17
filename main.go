package worker_store

type StoreType int64

const (
	Memory StoreType = 0
	Etcd             = 1
	Redis            = 2
)

type WorkerStoreType struct {
	WorkerStoreType StoreType `json:"storeType"`
}

func InitWorkerStore(storeType StoreType) WorkerStoreType {
	return WorkerStoreType{
		WorkerStoreType: storeType,
	}
}
