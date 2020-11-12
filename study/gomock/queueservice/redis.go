package queueservice

import "errors"

type KafkaService struct {
	message string
}

func (s *KafkaService) Push(message string) (err error) {
	println("Sending message to redis: " + message)
	s.message = message
	return
}

func (s *KafkaService) Pop() (message string, err error) {
	if s.message == "" {
		err = errors.New("err empty queue")
		return
	}
	println("Receiving message from redis: " + s.message)
	return
}
