package worker_store

type WorkerStoreJobTypeEnum int64

const (
	File WorkerStoreJobTypeEnum = 0
)

type WorkerStoreJob struct {
	Id      uint64                `json:"id"`
	Type    WorkerStoreType       `json:"type"`
	Payload WorkerStoreJobPayload `json:"payload"`
}

type WorkerStoreJobType struct {
	Type WorkerStoreJobTypeEnum `json:"type"`
}

type WorkerStoreJobPayload struct{}
