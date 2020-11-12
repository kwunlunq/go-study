package queueservice

type QueueService interface {
	Push(message string) (err error)
	Pop() (message string, err error)
}
